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
    - create help message
    - expand add/delete to accept infinitely many arguments
    - make parser retain the file's indentation
    - handle errors more gracefully. Abstract a function, make error codes
    - add array length checking if JSON array length exceeds array capacity
    - add ability to set index for list (initial is 0-indexed, but can set it to 1)

# NOTES:
    - Marshal functions are for buffers. NewEncoder functions do not. If you want to read JSON,
    manipulate it, and send it back, you need to use marshal.
    - Any time you process a buffer, you need to know the amount of bytes actually read into it. This means 
    that every time you use that buffer for something after you pass data into it, you need to specify
    the range of characters in the buffer that actually contain the data that you care about, by using a 
    slice operator. Typical usage looks like passing in "buffer[:amntBytes]"
    - JSON files might have to end in a newline, but I'm unsure. Additionally parsing JSON ruins
    its indentation (but there's functions to fix that)
