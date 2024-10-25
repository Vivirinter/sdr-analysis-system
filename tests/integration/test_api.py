import requests

def test_api():
    response = requests.get("http://localhost:8080/data")
    assert response.status_code == 200
    assert response.text == "SDR Data"