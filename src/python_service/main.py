from concurrent import futures
import logging
from queue import Empty
from urllib import response
import grpc
import python_service_list_pb2_grpc as genpb2
import python_service_list_pb2
from grpc_reflection.v1alpha import reflection
from pymongo import MongoClient


class EventsServicer(genpb2.EventsServicer):
    # Define our grpc event -> ListOneItem
    def ListOneItem(self, request, context):
        # Get items in itemsCollection
        item_details = collection_name.find()
        for item in item_details:
          print(item)
          print("\n")
        return python_service_list_pb2.ResponseValue(response = "STRINGG")

# Defining database components
def get_database():
  # Provide the mongodb url
  CONNECTION_STRING = "mongodb://citizix:S3cret@localhost:27017"
  client = MongoClient(CONNECTION_STRING)

  # Create the database for our microservice practice
  return client['itemsDB']

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  genpb2.add_EventsServicer_to_server(
  EventsServicer(), server)
  SERVICE_NAMES = (
    python_service_list_pb2.DESCRIPTOR.services_by_name['Events'].full_name,
    reflection.SERVICE_NAME,
  )   
  reflection.enable_server_reflection(SERVICE_NAMES, server)
  server.add_insecure_port('[::]:50051')
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
    dbname = get_database()
    collection_name = dbname["itemsCollection"]
    

    logging.basicConfig()
    print("Python gRPC server serve..")
    serve()

