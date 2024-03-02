import json
from dataclasses import dataclass, field, asdict
from typing import Optional, List, Any

import requests

from examples.mockingbird.mockingbooks_py.api_setup import get_headers, api_v1_path


@dataclass
class EntitiesFilter:
    nickname: str
    platform: str
    firstName: Optional[str] = None
    lastName: Optional[str] = None
    labels: List[str] = field(default_factory=list)
    metadataJsonb: Optional[Any] = None  # Use Any, could also use dict or a more specific type based on your needs
    metadataText: Optional[str] = None
    sinceUnixTimestamp: Optional[int] = None

    def to_dict(self) -> dict:
        data = asdict(self)
        # Correcting the field name for JSON serialization
        if 'sinceUnixTimestamp' in data:
            data['sinceTimestampUnix'] = data.pop('sinceUnixTimestamp')
        # Optional handling for metadataJsonb to ensure it's serialized properly
        if 'metadataJsonb' in data and data['metadataJsonb'] is not None:
            data['metadataJsonb'] = json.dumps(data['metadataJsonb'])
        return {k: v for k, v in data.items() if v is not None}


def create_entity(entity: EntitiesFilter):
    url = api_v1_path + "/entities/ai"  # Adjusted endpoint for entities creation
    headers = get_headers()  # Assuming get_headers is defined elsewhere

    # Serialize the entity to a dict and then to JSON
    entity_data = entity.to_dict()
    pretty_data = json.dumps(entity_data, indent=4)
    print(pretty_data)

    resp = requests.post(url, json=entity_data, headers=headers)
    # Check the response status
    if resp.status_code == 200:
        print("Entity created successfully!")
        return resp.json()  # Return the JSON response on success
    else:
        print("Failed to create entity. Status Code:", resp.status_code)
        return resp.json()  # Still return the JSON response which might contain error details


def get_entities(entity: EntitiesFilter):
    url = api_v1_path + "/entities/ai"  # Adjusted endpoint for entities creation
    headers = get_headers()  # Assuming get_headers is defined elsewhere

    # Serialize the entity to a dict and then to JSON
    entity_data = entity.to_dict()
    pretty_data = json.dumps(entity_data, indent=4)
    print(pretty_data)

    resp = requests.get(url, json=entity_data, headers=headers)
    # Check the response status
    if resp.status_code == 200:
        print("Entity retrieved successfully!")
        return resp.json()  # Return the JSON response on success
    else:
        print("Failed to get entity. Status Code:", resp.status_code)
        return resp.json()  # Still return the JSON response which might contain error details


if __name__ == "__main__":
    # Sample data for an EntitiesFilter instance
    sample_entity = EntitiesFilter(
        nickname="SampleDemo",
        platform="PlatformSample",
        labels=["label1", "label2"],
        sinceUnixTimestamp=1234567890
    )

    print(get_entities(sample_entity))

    pretty_data1 = get_entities(sample_entity)
    pretty_data2 = json.dumps(pretty_data1, indent=4)
    print(pretty_data2)

    # print(json.dumps(sample_entity.to_dict()))
    # Create the entity using the dummy function

    # response = create_entity(sample_entity)
    # print(response)
