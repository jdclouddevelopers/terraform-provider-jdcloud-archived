git filter-branch -f --env-filter '
if [ "$GIT_AUTHOR_NAME" = "liangxiaohan" ]
then
export GIT_AUTHOR_NAME="liangxiaohan"
export GIT_AUTHOR_EMAIL="liangxiaohan95@gmail.com"
fi
' HEAD
 
git filter-branch -f --env-filter '
if [ "$GIT_COMMITTER_NAME" = "liangxiaohan" ]
then
export GIT_COMMITTER_NAME="liangxiaohan"
export GIT_COMMITTER_EMAIL="liangxiaohan95@gmail.com"
fi
' HEAD
