from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework import status
from random import randint


import time
import random
import requests

from concurrent import futures

# Здесь укажите URL основного сервиса, где находится ваш эндпойнт
CALLBACK_URL = "http://127.0.0.1:8888/UpdateFlightAsyncResult/"
access_hash = "ASLDKjalksdjalskjdlk12lk3jfjkfdsfdasdASIODU*As"
executor = futures.ThreadPoolExecutor(max_workers=1)


def get_random_async_status(flight_id):
    time.sleep(4)
    tmp = randint(1,10)
    if tmp <= 8:
        return {
            "flight_id": flight_id,
            "result": "успешный полёт"
        }
    elif tmp >8:
        return {
            "flight_id": flight_id,
            "result": "АМС разбилась"
        }

def async_status_callback(task):
    try:
        result = task.result()
        print("Sending result update:", result)
    except futures._base.CancelledError:
        return
    except Exception as e:
        print("Failed to get a result:", e)
        return

    nurl = CALLBACK_URL + str(result["flight_id"])
    headers = {'Content-Type': 'application/json'}
    answer = {"result": result["result"],"access_hash":access_hash}



    # Здесь используем requests.put для отправки обновления статуса
    try:
        response = requests.put(nurl, json=answer, headers=headers, timeout=3)
        response.raise_for_status()
    except requests.RequestException as e:
        print("Failed to send the update:", e)


@api_view(['POST'])
def start_async_update(request):
    flight_id = request.data.get("flight_id")

    if flight_id:
        task = executor.submit(get_random_async_status, flight_id)
        task.add_done_callback(async_status_callback)
        return Response({"message": "Update started"}, status=status.HTTP_202_ACCEPTED)
    return Response({"error": "flight_id is missing"}, status=status.HTTP_400_BAD_REQUEST)