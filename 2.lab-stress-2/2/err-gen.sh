number=`shuf -i 0-4 -n 1`

# If <=1 success else error
if [ $number -le 1 ]; then
  echo "1"
  exit 0
else
  echo "$(shuf -i 2-5 -n 1)"
  exit 1
fi
