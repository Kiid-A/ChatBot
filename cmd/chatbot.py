from flask import Flask, request, jsonify
import yaml
import re
import pymysql

app = Flask(__name__)

# 配置数据库连接
DB_CONFIG = {
    "host": "localhost",   
    "port": 3306,          
    "user": "chatbot",     
    "password": "1234",    
    "database": "chat_db"  
}

def get_db_connection():
    """创建数据库连接"""
    return pymysql.connect(
        host=DB_CONFIG["host"],
        port=DB_CONFIG["port"],
        user=DB_CONFIG["user"],
        password=DB_CONFIG["password"],
        database=DB_CONFIG["database"],
        cursorclass=pymysql.cursors.DictCursor
    )

def query_user_balance(user_id):
    """查询用户余额"""
    connection = get_db_connection()
    try:
        with connection.cursor() as cursor:
            sql = "SELECT balance FROM users WHERE user_id = %s"
            cursor.execute(sql, (user_id,))
            result = cursor.fetchone()
            if result:
                return f"${result['balance']:.2f}"
            else:
                return "Balance not found"
    finally:
        connection.close()

def update_user_balance(user_id, amount):
    """更新用户余额（充值操作）"""
    connection = get_db_connection()
    try:
        with connection.cursor() as cursor:
            sql = "UPDATE users SET balance = balance + %s WHERE user_id = %s"
            cursor.execute(sql, (amount, user_id))
            connection.commit()
            return True
    except Exception as e:
        return False
    finally:
        connection.close()


class StateMachine:
    def __init__(self, states, transitions, initial_state):
        self.states = states
        self.transitions = transitions
        self.current_state = initial_state
        self.initial_state = initial_state  # 保存初始状态

    def reset_state(self):
        """重置状态为初始状态"""
        self.current_state = self.initial_state

    def handle_input(self, user_input, context):
        """根据用户输入和当前状态处理状态机的转换"""
        for transition in self.transitions:
            if transition["from"] == self.current_state:
                # 匹配输入条件
                if transition["input"] == "*" or re.match(transition["input"], user_input):
                    # 动态处理响应内容
                    response = transition["response"]
                    if "{balance}" in response:
                        # 余额查询时触发数据库查询
                        response = response.replace("{balance}", context.get_balance())
                    if "{amount}" in response:
                        # 充值时更新余额并修改响应
                        context.set_amount(user_input)
                        amount = float(user_input)  # 充值金额
                        update_user_balance(context.user_data["user"], amount)  # 更新数据库
                        response = response.replace("{amount}", user_input)
                    # 状态转移
                    self.current_state = transition["to"]
                    return response
        # 如果没有找到匹配的转换，返回默认信息
        self.reset_state()
        return "I don't understand. You can type 'help' to check the menu."


class Context:
    """负责存储用户上下文信息及业务逻辑"""
    def __init__(self, user, get_balance_func):
        self.user_data = {"user": user}
        self.get_balance_func = get_balance_func
        self.balance = None  # 缓存余额

    def get_balance(self):
        """返回用户的余额"""
        if not self.balance:  # 如果余额缓存为空，查询数据库
            self.balance = self.get_balance_func(self.user_data["user"])
        return self.balance

    def set_amount(self, amount):
        """设置用户充值金额"""
        self.user_data["amount"] = amount

    def get_amount(self):
        """获取充值金额"""
        return self.user_data.get("amount", "0")


def load_dsl(file_path):
    """加载DSL文件"""
    with open(file_path, "r", encoding="utf-8") as file:
        dsl = yaml.safe_load(file)
    return dsl


def create_state_machine(dsl):
    """根据DSL配置生成状态机"""
    states = dsl.get("states", [])
    transitions = dsl.get("transitions", [])
    initial_state = states[0] if states else None
    return StateMachine(states, transitions, initial_state)


@app.route('/newchat/<user_id>', methods=['POST'])
def new_chat(user_id):
    """新聊天接口，传入user_id，初始化状态机"""
    # 加载DSL文件并创建状态机
    dsl_file_path = "cmd/dialogue_dsl.yaml"
    dsl = load_dsl(dsl_file_path)
    chatbot = create_state_machine(dsl)

    # 创建上下文并传入用户ID
    context = Context(user_id, query_user_balance)

    # 返回初始状态的响应
    response = chatbot.handle_input("start", context)
    
    return jsonify({"response": response})


@app.route('/chat/<user_id>/<chat_id>', methods=['POST'])
def chat(user_id, chat_id):
    """处理用户输入并返回状态机响应"""
    # 创建上下文并传入用户ID
    context = Context(user_id, query_user_balance)

    # 获取用户输入
    user_input = request.json.get('input', '')
    
    # 加载DSL文件并创建状态机
    dsl_file_path = "cmd/dialogue_dsl.yaml"
    dsl = load_dsl(dsl_file_path)
    chatbot = create_state_machine(dsl)

    # 获取当前状态的响应
    response = chatbot.handle_input(user_input, context)
    
    return jsonify({"response": response})


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)
