#include <linux/kernel.h>
#include <linux/module.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("931ce5c4a04d");

int init_module(void)
{
	pr_debug("Hello world!");
	return 0;
}

void cleanup_module(void)
{
	pr_debug("Peace out, world!");
}

module_init(init_module);
module_exit(cleanup_module);
