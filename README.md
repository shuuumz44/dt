*Go CLI Task Manager*

Exactly what it sounds like.

# USAGE:
	-- add task
	task add "go on a walk"
	
	-- update task
	task update 14 "drive to school"

    -- delete task
	task delete 14

	-- mark to do / in progress / complete
	task mark 0 14
	
	-- list all tasks
	task list

	-- list by status
	task list todo
	task list doing
	task list done

# FEATURES:
	@ adding 
		-- unique key generation
		-- keys in order, 1-indexed
		-- numbers reorder when a value earlier
		than the last is completed (structure is a linked list)

	@ updating
        -- changes the name of the specified task. If no replacement name is specified, update the task's time.

	@ deleting
        (correct indexes after list mutation)

	@ marking
        -- status listed as an integer:
            0 -> todo
            1 -> doing
            2 -> done

	@ listing

to do:
    - use JSON parsing from concept
    - properly set up linked list/manage pointers
    - fix scroll() segfault
    - expand add/delete to accept infinitely many arguments
    - make parser retain the file's indentation

# NOTES:
    - Marshal functions are for buffers. NewEncoder functions do not. If you want to read JSON,
    manipulate it, and send it back, you need to use marshal.
    - JSON may or may not be understood in practice to end in a newline. Additionally, parsing JSON
    destroys the file's indentation. This is barely a warning however because 99% of the
    time JSON files are meant to be generated, not typed by hand. Also JSON is explicitly understood
    to be a stream of bytes, and all files end in EOF, and...anyway.
