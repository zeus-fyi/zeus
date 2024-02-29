import json  # Import the json module

import requests

from examples.mockingbird.mockingbooks_py.api_setup import get_headers, api_v1_path


def get_task(tid):
    url = api_v1_path + "/task/ai/" + tid
    headers = get_headers()
    response = requests.get(url, headers=headers)

    if response.status_code == 200:
        task_data = response.json()
        pretty_task_data = json.dumps(task_data, indent=4)
        with open('task-' + tid + '.json', 'w') as json_file:
            json.dump(task_data, json_file, indent=4)
        print(pretty_task_data)
    else:
        print("Failed to fetch task data. Status Code:", response.status_code)
    return response.json()


def create_analysis_task(task):
    url = api_v1_path + "/tasks/ai"
    headers = get_headers()
    response = requests.post(url, json=task, headers=headers)
    # Check the response status
    if response.status_code == 200:
        print("Task created successfully!")
    else:
        print(response.json())
        print("Failed to create task. Status Code:", response.status_code)
    return response.json()


if __name__ == '__main__':
    with open('templates/analysis.json', 'r') as file:
        payload = json.load(file)

    payload['taskName'] = 'demo-analysis-task'
    payload['taskGroup'] = 'demo'
    payload['model'] = 'gpt-3.5-turbo-0125'
    payload['prompt'] = ('how do you think AI will respond to being offered a'
                         ' decision tree with cost assigned options and a budget?')
    print(payload)
    create_analysis_task(payload)
