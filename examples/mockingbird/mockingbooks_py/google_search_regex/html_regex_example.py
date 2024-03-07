import json
import re
import time

from examples.mockingbird.mockingbooks_py.entities import EntitiesFilter, search_entities
from examples.mockingbird.mockingbooks_py.google_search_regex.dynamic_google_search import start_wf

agg_prompt = ("Can you summarize the relevant person and/or associated business entity from the search results? It "
              "should use the most relevant search result that matches best and ignore others to prevent mixing"
              " multiple profiles. I want to know what this person does and what kind of role they perform, and "
              "summarize any other details that you can find and you should add more weight"
              "to LinkedIn information, you should strive for accuracy over greater inclusion."
              "You should also make what platform you select clear"
              "with the metadata sources that are associated from that platform, business, ie. LinkedIn, Twitter, etc."
              "so that we associate the correct entity metadata with the correct platforms.")


# Function to process each <li> element
def process_li_element(li_text):
    # Define the patterns for extracting person's name and company
    name_pattern = r'data-anonymize="person-name">\s*([^<]+?)\s*</span>'
    # Updated company pattern to stop at the first '<' character, ensuring no '>' is included
    company_pattern = r'data-anonymize="company-name"[^>]*>\s*([^<]+?)\s*<'

    # Search for name and company within the <li> element text
    name_match = re.search(name_pattern, li_text)
    company_match = re.search(company_pattern, li_text)

    # If both name and company are found, return them
    if name_match and company_match:
        return name_match.group(1), company_match.group(1)
    else:
        return None


def iterate_on_matches():
    # List to store all matches
    all_matches = []
    count = 0

    # Process files from page1 to page60
    for page_num in range(1, 61):

        file_path = f'tmp/page{page_num}.txt'

        # Open the file and read its contents
        with open(file_path, 'r') as file:
            text = file.read()
        # Assuming 'text' contains the HTML content from your file
        # Define the regex pattern to capture each <li> element
        li_pattern = r'<li class="artdeco-list__item[^>]*>.*?</li>'

        # Find all <li> elements
        li_elements = re.findall(li_pattern, text, re.DOTALL)

        # Process each <li> element to extract names and companies
        matches = [process_li_element(li) for li in li_elements]

        # Filter out any None results
        matches = [match for match in matches if match is not None]

        # Iterate and print each match
        for name, company in matches:
            count += 1

            if count > 200:
                nc = f'Name: {name}, Company: {company}'
                print(nc)
                start_wf(nc, agg_prompt)
                all_matches += [(name, company)]
                count += 1
                time.sleep(10)
        print(count)


if __name__ == '__main__':
    iterate_on_matches()
    search_entities_f = EntitiesFilter()

    pretty_data1 = search_entities(search_entities_f)
    pretty_data2 = json.dumps(pretty_data1, indent=4)
    print(pretty_data2)

    for v in pretty_data1:
        print(v)
        print('---')
    # poll_run('1709446959958934000')
