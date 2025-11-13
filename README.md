# Maps Weather

## Execution
``.\maps-weather.exe -city={city} -api_key={your_api_key_here}``

## Params

- ``-city`` The city for which the check will be performed. For cities which are made up of more than one word (f.e San Francisco) please connect the words with a hyphen '-' as shown in the example below.
- ``-api-key`` The api-key for the OpenWeather service.

## Example Usage and Output
````
.\maps-weather.exe -city=San-Francisco -api_key=4bc0as420fb163115d721b21813g4d7g
````
#### Output:
```
COUNTRY: US
State: California
CITY: San-Francisco
OVERALL_WEATHER: moderate rain
TEMPERATURE: 14.82 C
FEELS_LIKE: 14.63 C
```

## Building
- ``go build`` will apply any code changes to the ``maps-weather.exe`` executable
- ``go help build`` for more info 

## Prerequisites
- Have an OpenWeather API key which can be gotten for free in 2 minutes from [openweathermap](https://openweathermap.org)