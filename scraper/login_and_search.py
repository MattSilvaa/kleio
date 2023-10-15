import os
import time

from bs4 import BeautifulSoup
from dotenv import load_dotenv
from selenium import webdriver
from selenium.common.exceptions import TimeoutException
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.chrome.service import Service

service = Service()
options = webdriver.ChromeOptions()
driver = webdriver.Chrome(service=service, options=options)

# Load environment variables from .env file
load_dotenv()

# Retrieve environment variables
LINKEDIN_USER = os.getenv('LINKEDIN_USER')
LINKEDIN_PW = os.getenv('LINKEDIN_PW')

# Initialize Selenium WebDriver

try:
    # Navigate to LinkedIn Login Page
    driver.get("https://www.linkedin.com/login")

    # Enter Username
    username_element = driver.find_element('id', 'username')
    username_element.send_keys(LINKEDIN_USER)

    # Enter Password
    password_element = driver.find_element('id', 'password')
    password_element.send_keys(LINKEDIN_PW)

    # Submit Login Form
    password_element.send_keys(Keys.RETURN)

    # Pause to allow login to complete
    time.sleep(3)

    driver.get(
        "https://www.linkedin.com/jobs/search/?currentJobId=3694175192&distance=25&geoId=104116203&keywords=Software%20Engineer&location=Seattle%20Washington")

    try:
        # Wait up to 10 seconds before throwing a TimeoutException
        # unless the condition (element is present) is satisfied
        element_present = EC.presence_of_element_located((By.CLASS_NAME, 'jobs-search-results-list__subtitle'))
        WebDriverWait(driver, 1000).until(element_present)

        # Now that we know the element is present, extract the page source
        page_source = driver.page_source

        # Parse page source with BeautifulSoup
        parsed_page_source = BeautifulSoup(page_source, 'html.parser')

        # Interact with the element
        job_count_element = parsed_page_source.find(class_='jobs-search-results-list__subtitle')
        if job_count_element:
            job_count = job_count_element.get_text(strip=True)
            print(f"Job count: {job_count}")
        else:
            print("Job count element not found")

    except TimeoutException:
        print("Timed out waiting for element to load")

finally:
    # Ensure the driver quits, even if there's an error
    driver.quit()
