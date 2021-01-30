Makefile OpenSSL Example
---
```
# ------------------------------------------------------------------------
# PROJECT_CFG is defined as a default in the calling Makefile or passed
# from the command line: make PROJECT_CFG=brutus-1.0
#------------------------------------------------------------------------
#------------------------------------------------------------------------
# Define Default Compile Variant
#------------------------------------------------------------------------
COMPILE_VARIANT          = debug

#------------------------------------------------------------------------
# Define Default Kernel Link Variant
#------------------------------------------------------------------------
LINK_VARIANT             = static

#------------------------------------------------------------------------
# Default Target Platform Variant
#------------------------------------------------------------------------
TARGET_CPU               = ia32
TARGET_CPU_PLATFORM      = $(TARGET_CPU)

TARGET_OS                = Linux
TARGET_OS_VERSION        = generic
TARGET_OS_PLATFORM       = $(TARGET_OS)-$(TARGET_OS_VERSION)

TARGET_PLATFORM          = $(TARGET_CPU_PLATFORM)_$(TARGET_OS_PLATFORM)
TARGET_PLATFORM_BASE     = $(TARGET_CPU)_$(TARGET_OS)

#------------------------------------------------------------------------
# Define Default Host Platform
#------------------------------------------------------------------------
HOST_CPU                 = ia32
HOST_CPU_PLATFORM        = $(HOST_CPU)

HOST_OS                  = Linux
HOST_OS_VERSION          = generic
HOST_OS_PLATFORM         = $(HOST_OS)-$(HOST_OS_VERSION)

#------------------------------------------------------------------------
# Toolchain Variants
#------------------------------------------------------------------------
TOOLCHAIN                = mvista

ifeq ($(TOOLCHAIN), mvista)

  TOOLCHAIN_VERSION      = i2.02
  TOOLCHAIN_ROOT         = /auto/ide/$(TOOLCHAIN_VERSION)/hardhat/devkit

  TOOLCHAIN_HOME_ia32    = $(TOOLCHAIN_ROOT)/x86/pentium4
  TOOLCHAIN_HOME_mips    = $(TOOLCHAIN_ROOT)/mips/fp_be
  TOOLCHAIN_HOME         = $(TOOLCHAIN_HOME_$(TARGET_CPU))

  TOOLCHAIN_PATH         = $(TOOLCHAIN_HOME)/bin
  TOOLCHAIN_TARGET       = $(TOOLCHAIN_HOME)/target

  TOOLCHAIN_PREFIX_ia32  = pentium4-
  TOOLCHAIN_PREFIX_mips  = mips_fp_be-
  TOOLCHAIN_PREFIX       = $(TOOLCHAIN_PREFIX_$(TARGET_CPU))
  TOOLCHAIN_PLATFORM     = $(TOOLCHAIN)-$(TOOLCHAIN_VERSION)

else
  TOOLCHAIN_VERSION      =
  TOOLCHAIN_ROOT         = /usr

  TOOLCHAIN_HOME_ia32    = $(TOOLCHAIN_ROOT)
  TOOLCHAIN_HOME_mips    = $(TOOLCHAIN_ROOT)
  TOOLCHAIN_HOME         = $(TOOLCHAIN_HOME_$(TARGET_CPU))

  TOOLCHAIN_PATH         = $(TOOLCHAIN_HOME)/bin
  TOOLCHAIN_TARGET       =

  TOOLCHAIN_PREFIX_ia32  =
  TOOLCHAIN_PREFIX_mips  =
  TOOLCHAIN_PREFIX       = $(TOOLCHAIN_PREFIX_$(TARGET_CPU))
  TOOLCHAIN_PLATFORM     = $(TOOLCHAIN)
endif

CC                       = $(TOOLCHAIN_PATH)/$(TOOLCHAIN_PREFIX)gcc
AR                       = $(TOOLCHAIN_PATH)/$(TOOLCHAIN_PREFIX)ar -r
LD                       = $(TOOLCHAIN_PATH)/$(TOOLCHAIN_PREFIX)ld
RANLIB                   = $(TOOLCHAIN_PATH)/$(TOOLCHAIN_PREFIX)ranlib
LDLIB_PATH               = -L$(TOOLCHAIN_TARGET)/usr/lib
INCLUDES                 = -I$(TOOLCHAIN_TARGET)/usr/include

#------------------------------------------------------------------------
# Resolve TOOLCHAIN configuration
#------------------------------------------------------------------------
HOST_PLATFORM            = $(HOST_CPU_PLATFORM)_$(TOOLCHAIN_PLATFORM)_$(HOST_OS_PLATFORM)
HOST_PLATFORM_BASE       = $(HOST_CPU)_$(HOST_OS)_$(TOOLCHAIN)

#------------------------------------------------------------------------
# Define the Build Platform
#------------------------------------------------------------------------
BUILD_PLATFORM           = $(HOST_PLATFORM)_$(TARGET_PLATFORM)
BUILD_PLATFORM_BASE      = $(TOOLCHAIN_PLATFORM)_$(TARGET_PLATFORM_BASE)_$(COMPILE_VARIANT)-$(LINK_VARIANT)

#------------------------------------------------------------------------
# Default Targets
#------------------------------------------------------------------------
SHORT_ALIASES            = ia32
SHORT_ALIASES           += mips

ALIASES                  = $(SHORT_ALIASES:%=%/%)

DEFAULT_RELEASE_TARGETS  = $(SHORT_ALIASES)
DEFAULT_OPENSRC_TARGETS  = $(SHORT_ALIASES)

DEFAULT_TARGET           = all
ALL_TARGETS              = setup depend build install checksum
SETUP_TARGETS            =  build-dirs log-cfg graft fixes configure
NUKE_TARGETS             = graft-clean

#------------------------------------------------------------------------
# Define the deafult build destination directories
#------------------------------------------------------------------------
RESULTS_DIR              = $(PKG)_osbld
RESULTS_ROOT_DIR         = $(PKG_PARENT_DIR)/$(RESULTS_DIR)
RESULTS_DEST_DIR         = $(RESULTS_ROOT_DIR)/$(PROJECT_CFG)/$(BUILD_PLATFORM_BASE)

BUILD_DIR                = build
BUILD_ROOT_DIR           = $(RESULTS_DEST_DIR)/$(BUILD_DIR)
BUILD_DEST_DIR           = $(BUILD_ROOT_DIR)

BUILD_LOG_DIR            = $(BUILD_DEST_DIR)/fubar/logs

#------------------------------------------------------------------------
# Define the default install destination directories
#------------------------------------------------------------------------
INSTALL_DIR              = install
INSTALL_ROOT_DIR         = $(RESULTS_DEST_DIR)/$(INSTALL_DIR)
INSTALL_DEST_DIR         = $(INSTALL_ROOT_DIR)

#------------------------------------------------------------------------
# Define the a list of all build/install destination directories
#------------------------------------------------------------------------
BUILD_DIRS               = $(BUILD_DEST_DIR)
BUILD_DIRS              += $(BUILD_LOG_DIR)
BUILD_DIRS              += $(BUILD_INSTALL_DIR)

#------------------------------------------------------------------------
# Define the Default Locations to get system commands
#------------------------------------------------------------------------
HOST_BIN_DIR            = /bin/
HOST_USR_BIN_DIR        = /usr/bin/

CAT                     = $(HOST_BIN_DIR)cat
ECHO                    = $(HOST_BIN_DIR)echo
FIND                    = $(HOST_USR_BIN_DIR)find
MKDIR                   = $(HOST_BIN_DIR)mkdir
MD5SUM                  = $(HOST_USR_BIN_DIR)md5sum
PATCH                   = $(HOST_USR_BIN_DIR)patch
PERL                    = $(HOST_USR_BIN_DIR)perl
PWD                     = $(HOST_BIN_DIR)pwd
RM                      = $(HOST_BIN_DIR)rm
TAR                     = $(HOST_BIN_DIR)tar
TEE                     = $(HOST_USR_BIN_DIR)tee
TEST                    = $(HOST_USR_BIN_DIR)test
XARGS                   = $(HOST_USR_BIN_DIR)xargs


#------------------------------------------------------------------------
# Define the Default Locations to get build utilities
#------------------------------------------------------------------------
HOST_UTILS_DIR          = $(BUILD_UTILS)/bin/

GRAFT                   = $(HOST_UTILS_DIR)graft
GRAFT_EXCLUDE           = .graft-exclude


MAKE_DIR                = $(TEST) -d $(@) || $(MKDIR) -p $(@)
MAKE_DEST_DIR           = $(TEST) -d $(@D) || $(MKDIR) -p $(@D)
TEE_OUTPUT              = 2>&1 | $(TEE) $(BUILD_LOG_DIR)/$@.out
UNPACK                  = $(TAR) -zxvf

#------------------------------------------------------------------------
# Define Configure Varaints
#------------------------------------------------------------------------
CONFIGURE_CMD           = ./Configure
CONFIGURE_OPTS          = no-shared no-idea no-mdc2 no-rc5
CONFIGURE_TARGET_ia32   = linux-pentium
CONFIGURE_TARGET_mips   = linux-mips
CONFIGURE_TARGET        = $(CONFIGURE_TARGET_$(TARGET_CPU))
CONFIGURE_PREFIX        = --prefix=$(INSTALL_DEST_DIR)
CONFIGURE_LDLIB_PATH    = $(LDLIB_PATH)
CONFIGURE_INCLUDES      = $(INCLUDES)

CONFIGURE               = $(CONFIGURE_CMD)
CONFIGURE              += $(CONFIGURE_OPTS)
CONFIGURE              += $(CONFIGURE_TARGET)
CONFIGURE              += $(CONFIGURE_PREFIX)
CONFIGURE              += $(CONFIGURE_LDLIB_PATH)
CONFIGURE              += $(CONFIGURE_INCLUDES)

#------------------------------------------------------------------------
# FIX Targets
# ------------------------------------------------------------------------
# CONFIGURE passes in INCLUDES along with CFLAGS and we wind up with
# gcc -I$(HOST_TOOLCHAIN) -I.. -I../../include and what we need is
# gcc -I.. -I../../include -I$(HOST_TOOLCHAIN). So we need to swith
# the string in the Makefile from $CFLAGS $INCLUDE to $INCLUDE $CFLAGS
#
FIX_MAKE_1              = cd $(BUILD_DEST_DIR) &&
FIX_MAKE_1             += $(FIND) . -name Makefile | xargs
FIX_MAKE_1             += $(PERL) -spi -e 's@\$\(MAKEDEPEND\) --
FIX_MAKE_1             += \$\(CFLAG\) \$\(INCLUDE\)@\$\(MAKEDEPEND\) --
FIX_MAKE_1             += \$\(INCLUDE\) \$\(CFLAG\)@'

# Same as FIX_MAKE_1, but in this case its the INLCUDES variable, whereas
# in FIX_MAKE_1 its the INCLUDE variable
#
FIX_MAKE_2              = cd $(BUILD_DEST_DIR) &&
FIX_MAKE_2             += $(FIND) . -name Makefile | xargs
FIX_MAKE_2             += $(PERL) -spi -e 's@\$\(MAKEDEPEND\) --
FIX_MAKE_2             += \$\(CFLAG\) \$\(INCLUDES\)@\$\(MAKEDEPEND\) --
FIX_MAKE_2             += \$\(INCLUDES\) \$\(CFLAG\)@'

# Since we create links to the source tree, the Makefile fails when
# attempting to copy over existing files. So we patch the Makefile to
# use cp -f
#
FIX_MAKE_3             = cd $(BUILD_DEST_DIR)/crypto/evp &&
FIX_MAKE_3            += $(PATCH) < ../../fubar/patch.crypto.evp.Makefile

# Same as FIX_MAKE_3
#
FIX_MAKE_4             = cd $(BUILD_DEST_DIR)/fips/sha1 &&
FIX_MAKE_4            += $(PATCH) < ../../fubar/patch.fips.sha1.Makefile

# domd is a front-end to make depend... and it is broken when attempting
# to override the default compiler. So we patched it to work correctly
#
FIX_DOMD               = cd $(BUILD_DEST_DIR)/util &&
FIX_DOMD              += $(PATCH) < ../fubar/patch.util.domd

fix-make-1 : FIX       = $(FIX_MAKE_1)
fix-make-2 : FIX       = $(FIX_MAKE_2)
fix-make-3 : FIX       = $(FIX_MAKE_3)
fix-make-4 : FIX       = $(FIX_MAKE_4)
fix-domd   : FIX       = $(FIX_DOMD)

FIX_TARGETS            = fix-make-1
FIX_TARGETS           += fix-make-2
FIX_TARGETS           += fix-make-3
FIX_TARGETS           += fix-make-4
FIX_TARGETS           += fix-domd


#------------------------------------------------------------------------
# Define Help Targets
#------------------------------------------------------------------------
HELP_FILE             = Make-help.txt
HELP_TARGETS_FILE     = Make-help-targets.txt
HELP_PROJECT_FILE     = Make-help-project.txt

help         : HELP   = @$(CAT) $(HELP_FILE)
help-targets : HELP   = @$(CAT) $(HELP_TARGETS_FILE)
help-project : HELP   = @$(CAT) $(HELP_PROJECT_FILE)


HELP_TARGETS          = help
HELP_TARGETS         += help-targets
HELP_TARGETS         += help-project

#------------------------------------------------------------------------
# Define External Dependencies
#------------------------------------------------------------------------

#------------------------------------------------------------------------
# Define Target Rules
#------------------------------------------------------------------------
CONFIGURE_RULE         = cd $(BUILD_DEST_DIR) &&
CONFIGURE_RULE        += $(CONFIGURE) $(TEE_OUTPUT)

DEPEND_RULE            = cd $(BUILD_DEST_DIR) &&
DEPEND_RULE           += $(MAKE) $@ MAKEDEPPROG="$(CC)"  $(TEE_OUTPUT)

BUILD_RULE             = cd $(BUILD_DEST_DIR) &&
BUILD_RULE            += $(MAKE) CC="$(CC)" AR="$(AR)" RANLIB="$(RANLIB)" $(TEE_OUTPUT)

INSTALL_RULE           = cd $(BUILD_DEST_DIR) &&
INSTALL_RULE          += $(MAKE) $@ CC="$(CC)" AR="$(AR)" RANLIB="$(RANLIB)" $(TEE_OUTPUT)

CHECKSUM_RULE          = cd $(INSTALL_DEST_DIR) &&
CHECKSUM_RULE         += $(FIND) . -type f | $(XARGS) $(MD5SUM) $(TEE_OUTPUT)

SHORT_ALIAS_RULE       = $(MAKE) all TARGET_CPU=$@

ALIAS_RULE             = $(MAKE) $(@F) TARGET_CPU=$(@D)

RELEASE_RULE           = for i in $(DEFAULT_RELEASE_TARGETS); do
RELEASE_RULE          += $(MAKE) all TARGET_CPU=$$i; done

OPENSRC_RULE           = for i in $(DEFAULT_OPENSRC_TARGETS); do
OPENSRC_RULE          += $(MAKE) all TOOLCHAIN=gnu TARGET_CPU=$$i; done

GRAFT_RULE             = @cd ../; for i in `$(FIND) . -name SCCS`; do
GRAFT_RULE            += $(ECHO) SCCS > $$i/../$(GRAFT_EXCLUDE); done;
GRAFT_RULE            += echo BitKeeper >> ../$(GRAFT_EXCLUDE)

GRAFT_SRC_RULE         = @$(GRAFT) -ivt $(BUILD_DEST_DIR) $(PKG_SRC_DIR)
GRAFT_SRC_RULE        += $(TEE_OUTPUT)

GRAFT_CLEAN_RULE       = @$(ECHO) "Removing Graft Link Management Files";
GRAFT_CLEAN_RULE      += cd ../; $(FIND) . -name $(GRAFT_EXCLUDE) |
GRAFT_CLEAN_RULE      += $(XARGS) $(RM) -f;

CLEAN_RULE             = $(RM) -rf $(BUILD_DEST_DIR)


CLOBBER_RULE           = $(RM) -rf $(INSTALL_DEST_DIR)

NUKE_RULE              = @$(ECHO) "Nuke all build variants";
NUKE_RULE             += $(ECHO) $(RESULTS_ROOT_DIR) |$(XARGS) -t $(RM) -rf

HELP_RULE              = $(CAT) $(HELP)

#------------------------------------------------------------------------
#
#------------------------------------------------------------------------
SHOW_PROJECT_CFG  = ------ $(PROJECT_CFG) ------------\n
SHOW_PROJECT_CFG += TARGET_CPU              : $(TARGET_CPU)\n
SHOW_PROJECT_CFG += TARGET_OS               : $(TARGET_OS)\n
SHOW_PROJECT_CFG += TARGET_PLATFORM         : $(TARGET_PLATFORM)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += HOST_CPU                : $(HOST_CPU)\n
SHOW_PROJECT_CFG += HOST_OS                 : $(HOST_OS)\n
SHOW_PROJECT_CFG += HOST_PLATOFRM           : $(HOST_PLATFORM)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += HOST_BIN_DIR            : $(HOST_BIN_DIR)\n
SHOW_PROJECT_CFG += HOST_USR_BIN_DIR        : $(HOST_USR_BIN_DIR)\n
SHOW_PROJECT_CFG += HOST_UTILS_DIR          : $(HOST_UTILS_DIR)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += TOOLCHAIN               : $(TOOLCHAIN)\n
SHOW_PROJECT_CFG += TOOLCHAIN_VERSION       : $(TOOLCHAIN_VERSION)\n
SHOW_PROJECT_CFG += TOOLCHAIN_PLATFORM      : $(TOOLCHAIN_PLATFORM)\n
SHOW_PROJECT_CFG += TOOLCHAIN_PATH          : $(TOOLCHAIN_PATH)\n
SHOW_PROJECT_CFG += TOOLCHAIN_TARGET        : $(TOOLCHAIN_TARGET)\n
SHOW_PROJECT_CFG += CC                      : $(CC)\n
SHOW_PROJECT_CFG += AR                      : $(AR)\n
SHOW_PROJECT_CFG += LD                      : $(LD)\n
SHOW_PROJECT_CFG += RANLIB                  : $(RANLIB)\n
SHOW_PROJECT_CFG += LDLIB_PATH              : $(LDLIB_PATH)\n
SHOW_PROJECT_CFG += INCLUDES                : $(INCLUDES)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += SHORT_ALIASES           : $(SHORT_ALIASES)\n
SHOW_PROJECT_CFG += ALIASES                 : $(ALIASES)\n
SHOW_PROJECT_CFG += DEFAULT_RELEASE_TARGETS : $(DEFAULT_RELEASE_TARGETS)\n
SHOW_PROJECT_CFG += DEFAULT_OPENSRC_TARGETS : $(DEFAULT_OPENSRC_TARGETS)\n
SHOW_PROJECT_CFG += DEFAULT_TARGET          : $(DEFAULT_TARGET)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += RESULTS_ROOT_DIR        : $(RESULTS_ROOT_DIR)\n
SHOW_PROJECT_CFG += RESULTS_DEST_DIR        : $(RESULTS_DEST_DIR)\n
SHOW_PROJECT_CFG += BUILD_DEST_DIR          : $(BUILD_DEST_DIR)\n
SHOW_PROJECT_CFG += BUILD_LOG_DIR           : $(BUILD_LOG_DIR)\n
SHOW_PROJECT_CFG += INSTALL_DEST_DIR        : $(INSTALL_DEST_DIR)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += CONFIGURE_TARGET        : $(CONFIGURE_TARGET)\n
SHOW_PROJECT_CFG += CONFIGURE               : $(CONFIGURE)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += ------ $(PROJECT_CFG) ------------\n
```