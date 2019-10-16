local karel = {name: 'Karel', phone: '111222333', car: 'Tesla', mail: 'kaja@tesla.com'};
local jaruna = {name: 'Jaruna', phone: '444555666', car: 'Cabrio', mail: 'jarca123@gmail.com'};
local vilem = {name: 'Vilem', phone: '999999999', car: 'Favorit', mail: 'vilda@kvilda.cz'};

local people = [karel, jaruna, vilem];
{
  Lide: people,
  Customers:{ 
    ['Customer '+ c.name]: {
    pepik: 'yes'
    }
    for c in people
  } 
  OfferEmails:[
    {[c.mail]: 'Dear ' + c.name + ', our company will contact you on phone number ' + c.phone + ' to offer you an issurace for your '+ c.car + ' !!'}
    for c in people
  ] 
  // if customer name is in winner list add another text message to the mail
  CustomersWinner: [ 
    {[c.mail]: 'Dear ' + c.name + ', our company will contact you on phone number ' + c.phone + ' to offer you an issurace for your '+ c.car + ' !!',}
    + {['Winner '+ c.mail]: if c.name == 'Karel' then 'Dear ' + c.name + ' you won 3 bilion $ !!' else 'No winner'}
    for c in people
    ]
  CustomersWinner1: [ 
    {[c.mail]: 'Dear ' + c.name + ', our company will contact you on phone number ' + c.phone + ' to offer you an issurace for your '+ c.car + ' !!',}
    + {'Winner': if c.name == 'Karel' then 'Congratulations, you won 4 bilion $ !!' else 'You are not a winner'}
    for c in people
    ]
  CustomersWinner2: [ 
    {[c.mail]: 'Dear ' + c.name + ', our company will contact you on phone number ' + c.phone + ' to offer you an issurace for your '+ c.car + ' !!',}
    {[if c.name == 'Karel' then 'Winner']: 'Congratulations, you won 4 bilion $ !!'}
    for c in people
    ]
} 
