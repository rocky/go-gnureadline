#ifndef _COMPLETE_H_
#define _COMPLETE_H_

#include <stdio.h>
#include <readline/readline.h>
#include <stdlib.h> // for free()

char* _go_readline_strarray_at(char **strarray, int idx) ;
int _go_readline_strarray_len(char **strarray);
extern rl_completion_func_t *_go_attempted_completion_function;

#endif /* _COMPLETE_H_ */
