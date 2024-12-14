import requests
from bs4 import BeautifulSoup
import csv

MENS_URL = "https://kikokostadinov.com/shop/men"
page = requests.get(MENS_URL)

# print(page.text)
soup = BeautifulSoup(page.text, features="html.parser")
print(soup.title)

# clear file
with open('men-products.csv', 'w') as file:
    pass

fieldnames = ['Product', 'Link', 'Image']
data = []
products = soup.select('li.product.all.men')
with open('men-products.csv', 'w') as f:
    writer = csv.DictWriter(f, fieldnames=fieldnames)
    writer.writeheader()  # Write the header row
    for product in products:
        item = {}
        item['Product'] = product.select_one('figcaption').select_one('span').contents[0]
        # print(product.find('span', class_='price'))
        item['Link'] = product.find('a').get('href')
        item['Image'] = product.find('img').get('data-src')
        data.append(item)
    writer.writerows(data)
   
WOMENS_URL = "https://kikokostadinov.com/shop/women"
page = requests.get(WOMENS_URL)

# print(page.text)
soup = BeautifulSoup(page.text, features="html.parser")
print(soup.title)    
        
# clear file
with open('women-products.csv', 'w') as file:
    pass

fieldnames = ['Product', 'Link', 'Image']
data = []
products = soup.select('li.product.all.women')
with open('women-products.csv', 'w') as f:
    writer = csv.DictWriter(f, fieldnames=fieldnames)
    # writer.writeheader()  # Write the header row
    for product in products:
        item = {}
        item['Product'] = product.select_one('figcaption').select_one('span').contents[0]
        # print(product.find('span', class_='price'))
        item['Link'] = product.find('a').get('href')
        item['Image'] = product.find('img').get('data-src')
        data.append(item)
    writer.writerows(data)