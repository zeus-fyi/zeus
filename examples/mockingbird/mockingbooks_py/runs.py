import json  # Import the json module

import requests

from examples.mockingbird.mockingbooks_py.api_setup import get_headers, api_v1_path


def get_runs():
    url = api_v1_path + "/runs/ai"
    headers = get_headers()
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        data = response.json()
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)


def get_run(run_id):
    url = api_v1_path + "/run/ai/" + run_id
    headers = get_headers()
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        data = response.json()
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)


if __name__ == '__main__':
    # get_runs()
    get_run('1704069081079680000')
