Make Driver
---
```
1 #------------------------------------------------------------------------
2 # Package Identification
3 #------------------------------------------------------------------------
4 PKG              = $(PKG_NAME)-$(PKG_VERSION)
5 PKG_PARENT_DIR  := $(shell cd ../..; pwd)
6 PKG_SRC_DIR     := $(shell cd ..; pwd)
7 
8 #------------------------------------------------------------------------
9 # Define the default Project properties to build with
10 #------------------------------------------------------------------------
11 PROJECT_CFG      = reference-1.0
12 
13 #------------------------------------------------------------------------
14 # Include Project Specific Properties
15 #------------------------------------------------------------------------
16 include Makefile.$(PROJECT_CFG)
17 
18 SHOW_CFG         = $(ECHO) -e "
19 SHOW_CFG        += PROJECT_CFG           : $(PROJECT_CFG)\n
20 SHOW_CFG        += PKG                   : $(PKG)\n
21 SHOW_CFG        += $(SHOW_PROJECT_CFG)\n
22 SHOW_CFG        += "
23 
24 #------------------------------------------------------------------------
25 # Build Targets
26 #------------------------------------------------------------------------
27 default: $(DEFAULT_TARGET)
28 
29 all    : $(ALL_TARGETS)
30 
31 setup  : $(SETUP_TARGETS)
32 
33 configure:
34         $(CONFIGURE_RULE)
35 
36 depend:
37         $(DEPEND_RULE)
38 
39 build:
40         $(BUILD_RULE)
41 
42 install:
43         $(INSTALL_RULE)
44 
45 checksum:
46         $(CHECKSUM_RULE)
47 
48 release:
49         $(RELEASE_RULE)
50 
51 opensrc:
52         $(OPENSRC_RULE)
53 
54 $(SHORT_ALIASES):
55         $(SHORT_ALIAS_RULE)
56 
57 $(ALIASES):
58         $(ALIAS_RULE)
59 
60 #------------------------------------------------------------------------
61 # Graft Targets (soft links from build tree to source tree)
62 #------------------------------------------------------------------------
63 graft:  graft-exclude graft-src
64 
65 graft-exclude:
66         $(GRAFT_RULE)
67 
68 graft-src: $(BUILD_DEST_DIR)
69         $(GRAFT_SRC_RULE)
70 
71 graft-clean:
72         $(GRAFT_CLEAN_RULE)
73 
74 
75 #------------------------------------------------------------------------
76 # Build Cleanup Targets
77 #------------------------------------------------------------------------
78 clean:
79         $(CLEAN_RULE)
80 
81 clobber:
82         $(CLOBBER_RULE)
83 
84 nuke: $(NUKE_TARGETS)
85         $(NUKE_RULE)
86 
87 #------------------------------------------------------------------------
88 # Build and Install Root Directories
89 #------------------------------------------------------------------------
90 build-dirs: $(BUILD_DIRS)
91 
92 $(BUILD_DIRS):
93         @$(MAKE_DIR)
94 
95 #------------------------------------------------------------------------
96 # Build Help Targets
97 #------------------------------------------------------------------------
98 $(HELP_TARGETS):
99         $(HELP)
100         @$(ECHO)
101 
102 #------------------------------------------------------------------------
103 # Build Makefile Fixes
104 #------------------------------------------------------------------------
105 fixes:
106         for i in $(FIX_TARGETS); do $(MAKE) $$i; done
107 
108 $(FIX_TARGETS):
109         $(FIX) $(TEE_OUTPUT)
110         @$(ECHO)
111 
112 #------------------------------------------------------------------------
113 # Makefile Debug Targets
114 #------------------------------------------------------------------------
115 log-cfg:
116         @($(SHOW_CFG)) $(TEE_OUTPUT)
117 
118 show-cfg:
119         @$(SHOW_CFG
```