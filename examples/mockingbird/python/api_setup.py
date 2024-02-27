import yaml

config_file_path = '../../../test/configs/config.yaml'
with open(config_file_path, 'r') as file:
    config = yaml.safe_load(file)

bearer_token = config.get('BEARER', '')
headers = {
    'Authorization': f'Bearer {bearer_token}'
}

api_v1_path = "https://api.zeus.fyi/v1"
# api_v1_path = "http://localhost:9001/v1"


def get_headers():
    # Read the config file and extract the BEARER token
    with open(config_file_path, 'r') as fl:
        cfg = yaml.safe_load(fl)

    bt = cfg.get('BEARER', '')

    # Create and return the headers dictionary with the Authorization field
    hdrs = {
        'Authorization': f'Bearer {bt}'
    }
    return hdrs
