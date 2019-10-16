local Coffe(insugar=0, inmilk=false) = {
  // sugar = 0 => no sugar
  // sugar = 1 => one sugar brick
  // sugar = 2 => two sugar bricks
  // sugar = 3 => three sugar bricks

  // coffe without sugar and milk
  final_coffe: [
    {coffe: '50 ml',},
  ] + (
    if insugar > 3 then [{sugar: 3,},] else (if insugar > 0 then [{sugar: insugar,},] else [])
  )
  + (
    if inmilk then [{milk: 'yes',},] else []
  )
};

{
  Coffe: Coffe(),
  'Flat White': Coffe(0,true),
  'Flat White with 10 sugar blocks': Coffe(inmilk=true),
  'Sweet esspreso': Coffe(10,false),
  'Shitty esspreso': Coffe(-10,false),
}
