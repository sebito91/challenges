obj-m += hello.o

KERN_DIR=/lib/modules/$(shell uname -r)/build

all:
	$(MAKE) -C $(KERN_DIR) M=$(PWD) modules

clean:
	$(MAKE) -C $(KERN_DIR) M=$(PWD) clean
