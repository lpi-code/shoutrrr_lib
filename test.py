import ctypes
import os


shoutrrr_url = ""

lib = ctypes.CDLL(os.path.abspath("libshoutrrr.so"))
strbufurl = ctypes.c_char_p(shoutrrr_url.encode('utf-8'))
strbufmsg = ctypes.c_char_p("Hello World".encode('utf-8'))
lib.Send(strbufurl, strbufmsg)
