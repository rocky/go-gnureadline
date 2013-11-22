#include "complete.h" 
#include "_cgo_export.h"
#include <stdio.h>

char* _go_readline_strarray_at(char **strarray, int idx) 
{
  return strarray[idx];
}

int _go_readline_strarray_len(char **strarray)
{
  int sz = 0;
  while (strarray[sz] != NULL) {
    sz += 1;
  }
  return sz;
}

char **_attempted_completion_function(const char *text, int start, int end)
{
  char **c = goCallAttemptedCompletionFunction((char *)text, start, end);

  // readline will segfault if given a char *[1] = { NULL }
  if (c != NULL && c[0] == NULL) {
    free(c);
    c = NULL;
    rl_attempted_completion_over = 1;
  }
  return c;
}
rl_completion_func_t *_go_attempted_completion_function = _attempted_completion_function;
