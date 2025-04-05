import requests

url = "http://localhost:5555/store/clothes"

response = requests.get(url, headers={"api-key": "password"})
print(response.json())