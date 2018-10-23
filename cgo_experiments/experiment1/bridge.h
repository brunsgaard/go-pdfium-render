typedef struct message_s
{
  char message[255];
  int displayed;
} Message;

Message *create_message(char *msg);
void display_message(Message *message);
void free_message(Message *message);
