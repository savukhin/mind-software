import pytest

def test_ping():
    response = requests.get(API_URL + '/ping')
    assert response.status_code == 200
    assert response.json() == {'message': 'pong'}
    assert response.headers['content-type'] == 'application/json'

if __name__ == "__main__":
    pytest.main(["integrational.py"])