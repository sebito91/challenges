#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/usb.h>
#include <linux/hid.h>

static struct usb_device_id kb_id_table[] = {
	{ USB_INTERFACE_INFO(USB_INTERFACE_CLASS_HID,
			     USB_INTERFACE_SUBCLASS_BOOT,
			     USB_INTERFACE_PROTOCOL_KEYBOARD)
	},
	{ }  /* terminating entry */
};
MODULE_DEVICE_TABLE(usb, kb_id_table);

static int __init init_kbmodule(void)
{
	/* register this driver with USB subsystem */
	pr_debug("registered hello_kb_driver");

	return 0;
}

static void __exit exit_kbmodule(void)
{
	pr_debug("deregistered hello_kb_driver");
}

module_init(init_kbmodule);
module_exit(exit_kbmodule);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("931ce5c4a04d");
