#include <stdio.h>
#include <stdlib.h>

#include <fpdfview.h>


FPDF_DOCUMENT create_document(){
  FPDF_DOCUMENT doc;
  doc = (FPDF_DOCUMENT)malloc(sizeof(FPDF_DOCUMENT));
  doc = FPDF_LoadDocument("in.pdf", NULL);
  //FPDF_CloseDocument(doc);
  
  return doc;
}

void free_document(FPDF_DOCUMENT *doc)
{
  free(doc);
}

int main(int argc, char **argv)
{
  FPDF_InitLibrary();
  FPDF_DOCUMENT doc;
  FPDF_PAGE page;
  doc = FPDF_LoadDocument("in.pdf", NULL);
  if (doc == NULL)
  {
    printf("failed to open test document\n");
    return 1;
  }
  int numPages = FPDF_GetPageCount(doc);
  printf("document has %d pages\n", numPages);
  for (int i = 0; i < numPages; ++i)
  {
    page = FPDF_LoadPage(doc, i);
    if (page == NULL)
    {
      printf("failed to open page %d\n", i);
      continue;
      ;
    }
    double width = FPDF_GetPageWidth(page);
    double height = FPDF_GetPageHeight(page);
    printf("page %d is : %f x %f\n", i, width, height);
  }
  FPDF_DestroyLibrary();
  return 0;
}
