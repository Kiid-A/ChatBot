import requests

url = "http://localhost:5000/chat"
headers = {
    "Content-Type": "application/json"
}

inputs = [
    "help",
    "help",
    "continue",
    "check weather",
    "New York",
    "invalid input",
    "help",
    "exit"
]

for input_text in inputs:
    data = {
        "input": input_text
    }
    response = requests.post(url, json=data, headers=headers)
    print(f"Input: {input_text}")
    print("Status Code:", response.status_code)
    print("Response JSON:", response.json())
    print("-" * 40)
