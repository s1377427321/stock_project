#include <wchar.h>
#include <MediaInfoDLL/MediaInfoDLL.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <locale.h>
#include <limits.h>

const wchar_t *toWchar(const char *c)
{
    const size_t cSize = strlen(c)+1;
    wchar_t* wc = malloc(cSize * sizeof(wchar_t));
    mbstowcs (wc, c, cSize);
    return wc;
}

const char *toChar(const wchar_t *c)
{
    const size_t cSize = wcslen(c)*4+1;
    char* wc = malloc(cSize * sizeof(char));
    memset(wc, 0, cSize);
    wcstombs(wc, c, cSize);
    return wc;
}

void *GoMediaInfo_New() {
    return MediaInfo_New();
}

void GoMediaInfo_Delete(void *handle) {
    MediaInfo_Delete(handle);
}

size_t GoMediaInfo_OpenFile(void *handle, char *name) {
    return MediaInfo_Open(handle, toWchar(name));
}

size_t GoMediaInfo_OpenMemory(void *handle, char *bytes, size_t length) {
    MediaInfo_Open_Buffer_Init(handle, length, 0);
    MediaInfo_Open_Buffer_Continue(handle, bytes, length);

    return MediaInfo_Open_Buffer_Finalize(handle);
}

void GoMediaInfo_Close(void *handle) {
    MediaInfo_Close(handle);
}


const char *GoMediaInfoGetIdx(void *handle, int stream_index, int type, char *name) {
    return toChar(MediaInfo_Get(handle, type, stream_index,  toWchar(name), MediaInfo_Info_Text, MediaInfo_Info_Name));
}

const char *GoMediaInfoGet(void *handle, int type, char *name) {
    return GoMediaInfoGetIdx(handle, 0, type, name);
}

const char *GoMediaInfoOption(void *handle, char *name, char *value) {
    return toChar(MediaInfo_Option(handle, toWchar(name), toWchar(value)));
}

const char *GoMediaInfoInform(void *handle) {
    return toChar(MediaInfo_Inform(handle, 0));
}