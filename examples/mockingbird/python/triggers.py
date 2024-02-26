import json  # Import the json module

import requests

from examples.mockingbird.python.api_setup import get_headers, api_v1_path


def get_triggers():
    url = api_v1_path + "/actions/ai"
    headers = get_headers()
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        data = response.json()
        pretty_data = json.dumps(data, indent=4)
        print(pretty_data)
    else:
        print("Status Code:", response.status_code)


def create_or_update_trigger():
    pass


if __name__ == '__main__':
    with open('templates/triggers.json', 'r') as file:
        payload = json.load(file)
