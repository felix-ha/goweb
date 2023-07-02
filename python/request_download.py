import os
import requests

file_name = "file.txt"

url = "http://localhost:80/download"
json_content = {"filename": file_name}
response = requests.post(url, json=json_content)

if response.status_code == 200:
    directory_path = os.path.join(os.getcwd(), "downloads")
    os.makedirs(directory_path, exist_ok=True)
    file_path = os.path.join(directory_path, file_name)
    
    with open(file_path, "wb") as file:
        file.write(response.content)
