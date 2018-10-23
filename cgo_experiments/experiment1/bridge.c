#include "bridge.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

Message *create_message(char *msg)
{
  Message *message;

  message = (Message *)malloc(sizeof(Message));
  if (message != NULL)
  {
    strcpy(message->message, msg);
    message->displayed = 0;
  }

  return message;
}

void display_message(Message *message)
{
  printf("%s\n", message->message);
  message->displayed = 1;
}

void free_message(Message *message)
{
  free(message);
}
