local func(a,b) = {
  [if a then 'Ano' else 'Ne']: if b then 'ANO' else 'NE',
};
{
  obe_ano: func(true,true),
  obe_ne: func(false,false),
  b_ano: func(false,true),
  a_ano: func(true,false),
}
