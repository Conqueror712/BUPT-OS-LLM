import requests
import json
import time
import hashlib


def calculate_md5(input_string):
    md5 = hashlib.md5()
    md5.update(input_string.encode('utf-8'))
    encrypted = md5.hexdigest()
    return encrypted

def do_request():
    url = "https://api.baichuan-ai.com/v1/chat"
    api_key = "your_api_key"
    secret_key = "your_secret_key"

    data = {
        "model": "Baichuan2-53B",
        "messages": [
            {
                "role": "user",
                "content": "世界第一高峰是"
            }
        ]
    }

    json_data = json.dumps(data)
    time_stamp = int(time.time())
    signature = calculate_md5(secret_key + json_data + str(time_stamp))

    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + api_key,
        "X-BC-Request-Id": "your requestId",
        "X-BC-Timestamp": str(time_stamp),
        "X-BC-Signature": signature,
        "X-BC-Sign-Algo": "MD5",
    }

    response = requests.post(url, data=json_data, headers=headers)

    if response.status_code == 200:
        print("请求成功！")
        print("响应header:", response.headers)
        print("响应body:", response.text)
    else:
        print("请求失败，状态码:", response.status_code)


if __name__ == "__main__":
    do_request()


              