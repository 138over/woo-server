Makefile Targets Example
---
Replacement for suffix rules that do not support variants. 
```
#------------------------------------------------------------------------
# Object File Targets - C Language Source Files 
# Generated Object Targets must come first
#------------------------------------------------------------------------
$(OBJ_DEST_DIR)/%.dep: $(OBJ_DEST_DIR)/%.c
	$(C_DEP_RULE)

$(OBJ_DEST_DIR)/%.o: $(OBJ_DEST_DIR)/%.c
	$(C_O_RULE)

$(OBJ_DEST_DIR)/%.dep: $(SRC_DIR)/%.c
	$(C_DEP_RULE)

$(OBJ_DEST_DIR)/%.o: $(SRC_DIR)/%.c
	$(C_O_RULE)

$(OBJ_DEST_DIR)/%.i: $(OBJ_DEST_DIR)/%.c
	$(C_I_RULE)

$(OBJ_DEST_DIR)/%.i: $(SRC_DIR)/%.c
	$(C_I_RULE)

$(OBJ_DEST_DIR)/%.s: $(OBJ_DEST_DIR)/%.c
	$(C_S_RULE)

$(OBJ_DEST_DIR)/%.s: $(SRC_DIR)/%.c
	$(C_S_RULE)

#------------------------------------------------------------------------
# Shared Library Targets
#------------------------------------------------------------------------
$(SHARED_LIBRARY_TARGET): $(OBJSPIC)
	$(SHARED_LIBRARY_RULE)

#------------------------------------------------------------------------
# Static Library Targets
#------------------------------------------------------------------------
$(STATIC_LIBRARY_TARGET): $(OBJS)
	$(STATIC_LIBRARY_RULE)

#------------------------------------------------------------------------
# Program Target
#------------------------------------------------------------------------
$(PROGRAM_TARGET): $(LINK_LIBS) $(LINK_OBJS) $(OBJS)
	$(PROGRAM_RULE)
```
