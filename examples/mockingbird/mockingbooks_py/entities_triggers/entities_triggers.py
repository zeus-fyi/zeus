import json

from examples.mockingbird.mockingbooks_py.evals import create_or_update_eval
from examples.mockingbird.mockingbooks_py.retrievals import create_or_update_retrieval
from examples.mockingbird.mockingbooks_py.schemas import create_or_update_schema
from examples.mockingbird.mockingbooks_py.triggers import create_or_update_trigger


def io_create_entities():
    with open('mocks/io_create_entities.json', 'r') as file:
        ret_ce = json.load(file)
    data = create_or_update_retrieval(ret_ce)
    print(json.dumps(data, indent=4))


def io_search_entities():
    with open('mocks/io_search_entities.json', 'r') as file:
        ret_se = json.load(file)
    data = create_or_update_retrieval(ret_se)
    print(json.dumps(data, indent=4))


def create_entities_eval_fn():
    with open('mocks/eval_fn_create_entities.json', 'r') as file:
        eval_fn_pl = json.load(file)
    create_or_update_eval(eval_fn_pl)


def create_entities_trigger():
    with open('mocks/trigger_post_entities_data.json', 'r') as file:
        e_trig = json.load(file)
    create_or_update_trigger(e_trig)


def create_schemas():
    with open('mocks/schema_entities_filter.json', 'r') as file:
        schema_data = json.load(file)
    data = create_or_update_schema(schema_data)
    print(json.dumps(data, indent=4))

    with open('mocks/schema_entities_list.json', 'r') as file:
        schema_data = json.load(file)
    data = create_or_update_schema(schema_data)
    print(json.dumps(data, indent=4))


if __name__ == '__main__':
    create_schemas()
    pass
