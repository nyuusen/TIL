.PHONY:
	commit_date amend_commits

commit_date:
	./commit_with_date_and_push.sh

amend_commit:
	./amend_commit_and_push.sh