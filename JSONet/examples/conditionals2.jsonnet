local condition_test(value=1) = {
  karel: if value == 0 then 'nula' else ''
};
{
  'Karel je': condition_test(0),
  'Karel2 je': condition_test(),
}
