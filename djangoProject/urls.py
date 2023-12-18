
from django.contrib import admin
from django.urls import path
from app.views import start_async_update

urlpatterns = [
    path('start-async-update/',start_async_update, name='start-async-update'),
]
# asd