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
		than the last is completed
        -- structure is a linked list
        (you want all ids to self correct in a linked list. don't overcomplicate the obvious way to do this.)

	@ updating
        -- changes the name of the specified task. If no replacement name is specified, update the task's time.

	@ deleting
        -- correct indexes after list mutation

	@ marking
        -- status listed as an integer:
            0 -> todo
            1 -> doing
            2 -> done

	@ listing

to do:
    - actually add JSON & parsing
    - expand add/delete to accept infinitely many arguments

# NOTES:
    - Marshal/Unmarshal are meant for bytes/strings/chunks of data. Use encoder and decoder objects.
