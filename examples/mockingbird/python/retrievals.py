import json  # Import the json module

import requests

from examples.mockingbird.python.api_setup import get_headers

ret_id = ''

url = "https://api.zeus.fyi/v1/retrieval/ai/" + ret_id

payload = {}
headers = get_headers()
response = requests.get(url, headers=headers)

if response.status_code == 200:
    data = response.json()
    with open('ret-' + ret_id + '.json', 'w') as json_file:
        json.dump(data, json_file, indent=4)
    pretty_data = json.dumps(data, indent=4)
    print(pretty_data)
else:
    print("Failed to fetch task data. Status Code:", response.status_code)
