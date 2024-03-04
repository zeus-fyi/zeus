import json
import re

from examples.mockingbird.mockingbooks_py.entities import EntitiesFilter, search_entities

# Path to the file
file_path = 'tmp/tmp.txt'

# Open the file and read its contents
with open(file_path, 'r') as file:
    text = file.read()

# Define two patterns
pattern1 = r'data-anonymize="person-name">\s*([^<]+?)\s*</a>'
pattern2 = r'<span data-anonymize="company-name">\s*([^<]+?)\s*</span>'

agg_prompt = ("Can you summarize the relevant person and/or associated business entity from the search results? It "
              "should use the most relevant search result that matches best and ignore others to prevent mixing"
              " multiple profiles. I want to know what this person does and what kind of role they perform, and "
              "summarize any other details that you can find and you should add more weight"
              "to LinkedIn information, you should strive for accuracy over greater inclusion."
              "You should also make what platform you select clear"
              "with the metadata sources that are associated from that platform, business, ie. LinkedIn, Twitter, etc."
              "so that we associate the correct entity metadata with the correct platforms.")

# Find all matches for both patterns
matches1 = re.findall(pattern1, text)
matches2 = re.findall(pattern2, text)

# Check if lengths are equal
if len(matches1) > len(matches2):
    matches1, matches2 = matches2, matches1

offset_l = 80
offset_r = 100

# skip next
# for i in range(len(matches1)):
#     person_company = f"{i}:{matches1[i]} (company)"
#     if 0 + offset_l < i < 1 + offset_r:
#         print(person_company)
#         start_wf(person_company, agg_prompt)


# Example iteration and action simulation with 'start_wf' function.
# Since the 'start_wf' function is a conceptual example, we'll simulate its operation as a print statement.
# for person_company in formatted_people_companies:
#     # Simulate calling 'start_wf' function with the person_company tuple and aggregated prompt.
#     start_wf(person_company, agg_prompt)
#

# skip next

if __name__ == '__main__':
    search_entities_f = EntitiesFilter()

    pretty_data1 = search_entities(search_entities_f)
    pretty_data2 = json.dumps(pretty_data1, indent=4)
    print(pretty_data2)

    for v in pretty_data1:
        print(v)
        print('---')
    # poll_run('1709446959958934000')
