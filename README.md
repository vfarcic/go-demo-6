## Replace master with orig
xxx
```bash
git checkout orig

git merge -s ours master

git checkout master

git merge orig

git push
```
