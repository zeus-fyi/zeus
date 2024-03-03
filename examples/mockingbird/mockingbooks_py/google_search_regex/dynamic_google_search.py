import json

from examples.mockingbird.mockingbooks_py.entities import EntitiesFilter, search_entities
from examples.mockingbird.mockingbooks_py.runs import get_run
from examples.mockingbird.mockingbooks_py.workflows import start_or_schedule_wf


def start_wf(prompt=None, agg_prompt=None):
    with open('mocks/wf_exec.json', 'r') as file:
        wf_exec = json.load(file)

    # ex. use  import a csv file as inputs, and exec wf for each row
    # with open('search_entities/crm.csv', 'r') as file:
    #     inputs = json.load(file)
    # Inject override prompt, allows wf inputs to operate like function inputs

    tmp = {}
    if prompt:
        tmp['zeusfyi-verbatim'] = {'replacePrompt': prompt}
    if agg_prompt:
        if 'zeusfyi-verbatim' in tmp:  # Check if the key already exists
            # Merge 'agg_prompt' into 'zeusfyi-verbatim' if you need to combine them
            # For example, appending or modifying the existing entry
            # This is just an example and might need adjustment based on your actual needs
            tmp['zeusfyi-verbatim']['replacePrompt'] += f" {agg_prompt}"
        else:
            # If 'zeusfyi-verbatim' does not exist, create a new entry for 'agg_prompt'
            tmp['biz-lead-google-search-summary'] = {'replacePrompt': agg_prompt}

    if tmp:
        wf_exec['taskOverrides'] = tmp
    wf_item = {
        'workflowName': 'google-query-regex-index-wf',
    }
    wf_exec['workflows'] = [wf_item]
    pretty_data = json.dumps(wf_exec, indent=4)
    print(pretty_data)

    start_or_schedule_wf(wf_exec)


def poll_run(run_id):
    print(get_run(run_id))


def fetch_entities():
    search_entities_f = EntitiesFilter(

    )
    pretty_data1 = search_entities(search_entities_f)
    pretty_data2 = json.dumps(pretty_data1, indent=4)
    print(pretty_data2)


if __name__ == '__main__':
    start_wf()
    # fetch_entities()
    # poll_run('1709435054427175000')
