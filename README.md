elly
====

Old school PHP site constructor

The idea is that you can structure the project's PHP files any way you like. Each file would have a special header
that elly will use during a build process to place the file into the correct folder structure and with the correct 
file name. 

You could therefore keep controllers, models, and views grouped together by type or by function and not be concerned 
with the final structure. You can also easily modify the resulting folder structure without actually needing to remove
or create folders.

Under some serious construction

Things to be done, thought about, contemplated:
* Special comment block to declare a PHP file's location and other details
* Ensuring included and required files link up properly
