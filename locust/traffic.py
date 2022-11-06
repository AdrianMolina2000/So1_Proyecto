import json
from random import randrange
from locust import HttpUser, between, task
class readFile():
    def __init__(self):
        self.data = []

    def getData(self): #Metodo donde se obtiene un elemento de la lista de registros
        size = len(self.data) #Tamaño de los datos
        if size > 0:
            index = randrange(0, size) if size > 1 else 0
            return self.data[index]
        else:
            print("size -> 0")
            return None
    
    def loadFile(self):
        print("LOADING ...")
        try:
            with open("traffic.json", 'r') as file:
                self.data = json.loads(file.read())
        except Exception:
            print(f'Error : {Exception}')
class TrafficData(HttpUser):
    wait_time = between(1, 1) #Tiempo de espera entre registros
    reader = readFile()
    reader.loadFile()

    def on_start(self):
        print("On Start")
    
    @task
    def sendMessage(self):
        data = self.reader.getData() #Registro obtenido de la lista
        if data is not None:
            
            res = self.client.post("http://34.70.59.65.nip.io/input", json=data)
            response = res.json()
            print(response)
        else:
            print("Empty") #No hay mas datos por enviar
            self.stop(True)
    

     