from django.urls import path
from app.views import start_async_update

urlpatterns = [
    # Здесь определяем эндпойнт для начала асинхронного обновления
    path('start-async-update/', start_async_update, name='start-async-update'),
]