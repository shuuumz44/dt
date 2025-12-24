*Go CLI Task Manager*

Exactly what it sounds like.

USAGE:
	-- add task
	task add "go on a walk"
	
	-- update task
	task update 14 "drive to school"
	task delete 14

	-- mark in progress / complete
	task mark 14
	task done 14
	
	-- list all tasks
	task list

	-- list by status
	task list done
	task list todo
	task list mark


FEATURES:
	@ adding
		-- unique key generation
		-- keys in order, 1-indexed
		-- numbers reorder when a value earlier
		than the last is completed
        -- structure is a linked list

	@ updating

	@ deleting
        -- correct indexes after list mutation

	@ marking

	@ listing
