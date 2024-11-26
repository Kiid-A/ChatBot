from flask import Flask, request, jsonify
import yaml
import re

app = Flask(__name__)

class StateMachine:
    def __init__(self, states, transitions, initial_state):
        self.states = states
        self.transitions = transitions
        self.current_state = initial_state
        self.initial_state = initial_state  # 保存初始状态

    def handle_input(self, user_input, context):
        for transition in self.transitions:
            if transition["from"] == self.current_state:
                if transition["input"] == "*" or re.match(transition["input"], user_input):
                    # 动态处理业务逻辑
                    response = transition["response"]
                    # 替换动态数据
                    if "{balance}" in response:
                        response = response.replace("{balance}", context.get_balance())
                    if "{amount}" in response:
                        context.set_amount(user_input)  # 存储用户充值金额
                        response = response.replace("{amount}", user_input)
                    # 状态转移
                    self.current_state = transition["to"]
                    return response
        # 如果没有找到匹配的转换，返回提示信息
        return "I don't understand. You can type 'help' to check the menu."

    def reset_state(self):
        self.current_state = self.initial_state


class Context:
    """负责存储上下文和业务逻辑"""
    def __init__(self, user, get_balance_func):
        self.user_data = {"user": user}
        self.get_balance_func = get_balance_func

    def get_balance(self):
        # 调用外部传入的余额查询函数
        return self.get_balance_func(self.user_data["user"])

    def set_amount(self, amount):
        # 设置充值金额
        self.user_data["amount"] = amount

    def get_amount(self):
        # 获取充值金额
        return self.user_data.get("amount", "0")


def load_dsl(file_path):
    """加载DSL文件"""
    with open(file_path, "r", encoding="utf-8") as file:  # 指定编码为utf-8
        dsl = yaml.safe_load(file)
    return dsl


def create_state_machine(dsl):
    """根据DSL生成状态机"""
    states = dsl.get("states", [])
    transitions = dsl.get("transitions", [])
    initial_state = states[0] if states else None
    return StateMachine(states, transitions, initial_state)


# 示例外部调用的余额查询函数
def mock_get_balance(user):
    return f"${1000.00 + hash(user) % 1000:.2f}"  # 模拟返回不同用户的余额


# 加载DSL文件
dsl = load_dsl("./dialogue_dsl.yaml")

# 创建状态机
chatbot = create_state_machine(dsl)

# 上下文初始化
context = Context("test_user", mock_get_balance)

@app.route('/chat', methods=['POST'])
def chat():
    user_input = request.json.get('input', '')
    response = chatbot.handle_input(user_input, context)
    return jsonify({"response": response})

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)