struct STRUCT {
    const char* key;
    const char* value;
};

typedef int (*Callback)(int /*result*/, void* /*vpointer*/, const char* /*cstring*/, struct STRUCT* /*cstruct*/);

void set_callback(Callback callback);
int call_callback(int result, void* vpointer, const char* cstring);

