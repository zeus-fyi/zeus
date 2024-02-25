import json  # Import the json module

import requests

from examples.mockingbird.python.api_setup import get_headers

task_id = ''

url = "https://api.zeus.fyi/v1/task/ai/" + task_id

payload = {}
headers = get_headers()
response = requests.get(url, headers=headers)

if response.status_code == 200:
    task_data = response.json()
    pretty_task_data = json.dumps(task_data, indent=4)
    with open('task-' + task_id + '.json', 'w') as json_file:
        json.dump(task_data, json_file, indent=4)
    print(pretty_task_data)
else:
    print("Failed to fetch task data. Status Code:", response.status_code)
