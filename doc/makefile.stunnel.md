Makefile Stunnel Example
---
Top level interface to native build
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
CPP                      = $(TOOLCHAIN_PATH)/$(TOOLCHAIN_PREFIX)cpp
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
ALL_TARGETS              = setup build install checksum
SETUP_TARGETS            = build-dirs log-cfg graft fixes configure
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
HOST_UTILS_DIR          = /u4/pkramer/bin/

GRAFT                   = $(HOST_UTILS_DIR)graft
GRAFT_EXCLUDE           = .graft-exclude


MAKE_DIR                = $(TEST) -d $(@) || $(MKDIR) -p $(@)
MAKE_DEST_DIR           = $(TEST) -d $(@D) || $(MKDIR) -p $(@D)
TEE_OUTPUT              = 2>&1 | $(TEE) $(BUILD_LOG_DIR)/$@.out
UNPACK                  = $(TAR) -zxvf

#------------------------------------------------------------------------
# Defined External Dependencies
#------------------------------------------------------------------------
OPENSSL_PKG_NAME        = openssl
OPENSSL_PKG_VERSION     = 0.9.7g
OPENSSL_PKG             = $(OPENSSL_PKG_NAME)-$(OPENSSL_PKG_VERSION)_osbld
OPENSSL_PROJECT_CFG     = reference-1.0
OPENSSL_ROOT_DIR        = $(PKG_PARENT_DIR)/$(OPENSSL_PKG)
OPENSSL_LINK_VARIANT    = static
OPENSSL_INSTALL_DIR     = $(OPENSSL_ROOT_DIR)/$(OPENSSL_PROJECT_CFG)/$(BUILD_PLATFORM_BASE)/install
OPENSSL_INCLUDE_DIR     = $(OPENSSL_INSTALL_DIR)/include


#------------------------------------------------------------------------
# Define Configure Varaints
#------------------------------------------------------------------------
CONFIGURE_CMD           = ./configure
CONFIGURE_OPTS          =
CONFIGURE_TARGET_ia32   =
CONFIGURE_TARGET_mips   = --host=mips-linux --without-random
CONFIGURE_TARGET        = $(CONFIGURE_TARGET_$(TARGET_CPU))
CONFIGURE_PREFIX        = --prefix=$(INSTALL_DEST_DIR)
CONFIGURE_WITH          = --with-ssl=$(OPENSSL_INSTALL_DIR)
CONFIGURE_CC            = CC="$(CC)"
CONFIGURE_LD            = LD="$(LD)"
CONFIGURE_CPP           = CPP="$(CPP)"
CONFIGURE_LDLIB_PATH    = LDFLAGS="$(LDLIB_PATH)"
CONFIGURE_INCLUDES      = CPPFLAGS="-I$(OPENSSL_INCLUDE_DIR) $(INCLUDES)"

CONFIGURE               = $(CONFIGURE_CMD)
CONFIGURE              += $(CONFIGURE_OPTS)
CONFIGURE              += $(CONFIGURE_TARGET)
CONFIGURE              += $(CONFIGURE_PREFIX)
CONFIGURE              += $(CONFIGURE_WITH)
CONFIGURE              += $(CONFIGURE_CC)
CONFIGURE              += $(CONFIGURE_LD)
CONFIGURE              += $(CONFIGURE_CPP)
CONFIGURE              += $(CONFIGURE_LDLIB_PATH)
CONFIGURE              += $(CONFIGURE_INCLUDES)

#------------------------------------------------------------------------
# FIX Targets
# ------------------------------------------------------------------------
FIX_CONFIGURE_mips      = cd $(BUILD_DEST_DIR) &&
FIX_CONFIGURE_mips     += $(PATCH) < ./fubar/patch.configure
FIX_CONFIGURE           = $(FIX_CONFIGURE_$(TARGET_CPU))

FIX_STUNNEL_CNF         = cd $(BUILD_DEST_DIR)/tools   &&
FIX_STUNNEL_CNF        += $(PATCH) < ../fubar/patch.tools.stunnel.cnf

fix-stunnel-cnf : FIX   = $(FIX_STUNNEL_CNF)
fix-configure   : FIX   = $(FIX_CONFIGURE)

FIX_TARGETS             = fix-stunnel-cnf
FIX_TARGETS            += fix-configure

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

DEPEND_RULE            =

BUILD_RULE             = cd $(BUILD_DEST_DIR) &&
BUILD_RULE            += $(MAKE) CC="$(CC)" CPP="$(CPP)" AR="$(AR)"
BUILD_RULE            += LD="$(LD)" RANLIB="$(RANLIB)" $(TEE_OUTPUT)

INSTALL_RULE           = cd $(BUILD_DEST_DIR) &&
INSTALL_RULE          += $(MAKE) $@ CC="$(CC)" LD="$(LD)" AR="$(AR)"
INSTALL_RULE          += RANLIB="$(RANLIB)" $(TEE_OUTPUT)

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
SHOW_PROJECT_CFG += OPENSSL_PKG             : $(OPENSSL_PKG)\n
SHOW_PROJECT_CFG += OPENSSL_ROOT_DIR        : $(OPENSSL_ROOT_DIR)\n
SHOW_PROJECT_CFG += OPENSSL_INSTALL_DIR     : $(OPENSSL_INSTALL_DIR)\n
SHOW_PROJECT_CFG += OPENSSL_INCLUDE_DIR     : $(OPENSSL_INCLUDE_DIR)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += CONFIGURE_TARGET        : $(CONFIGURE_TARGET)\n
SHOW_PROJECT_CFG += CONFIGURE               : $(CONFIGURE)\n
SHOW_PROJECT_CFG += \n
SHOW_PROJECT_CFG += ------ $(PROJECT_CFG) ------------\n
```

