#include "bridge.h"
  
int main() {
  Message * msg = create_message("Hello, world");
  display_message(msg);
  free_message(msg);
  return 0;
}
