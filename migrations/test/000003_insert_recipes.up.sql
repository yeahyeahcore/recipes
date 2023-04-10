INSERT INTO recipes (
  name, 
  description, 
  photo_url, 
  ingredients, 
  steps, 
  total_steps_time, 
  updated_at, 
  created_at
) VALUES (
  'Chocolate Chip Cookies', 
  'Classic chocolate chip cookies made with butter, sugar, eggs, flour, and chocolate chips.', 
  'https://example.com/chocolate-chip-cookies.jpg', 
  '{butter, sugar, eggs, flour, chocolate chips}', 
  ARRAY[
    '{"title": "Mix the batter", "description": "In a large bowl, cream together the butter and sugar until light and fluffy. Beat in the eggs, one at a time, then stir in the flour and chocolate chips."}'::jsonb,
    '{"title": "Bake the cookies", "description": "Drop tablespoonfuls of dough onto a baking sheet and bake at 375 degrees F (190 degrees C) for 10 to 12 minutes or until lightly browned."}'::jsonb
  ], 
  '00:25:00', 
  NOW(), 
  NOW()
);

INSERT INTO recipes (
  name, 
  description, 
  photo_url, 
  ingredients, 
  steps, 
  total_steps_time, 
  updated_at, 
  created_at
) VALUES (
  'Spaghetti Carbonara', 
  'Traditional Italian pasta dish made with spaghetti, pancetta or bacon, eggs, and cheese.', 
  'https://example.com/spaghetti-carbonara.jpg', 
  '{spaghetti, pancetta or bacon, eggs, Parmesan cheese}', 
  ARRAY[
    '{
      "title": "Cook the spaghetti",
      "description": "Cook the spaghetti in a large pot of boiling salted water until al dente."
    }'::jsonb, 
    '{
      "title": "Cook the pancetta or bacon", 
      "description": "Cook the pancetta or bacon in a large skillet until crisp."
    }'::jsonb, 
    '{
      "title": "Beat the eggs and cheese", 
      "description": "In a bowl, whisk together the eggs and Parmesan cheese."
    }'::jsonb, 
    '{
      "title": "Combine the pasta and pancetta or bacon", 
      "description": "Drain the pasta and add it to the skillet with the pancetta or bacon. Toss well to combine."
    }'::jsonb, 
    '{
      "title": "Add the egg mixture", 
      "description": "Remove the skillet from the heat and add the egg mixture. Toss well to combine, being careful not to scramble the eggs."
    }'::jsonb
  ], 
  '00:20:00', 
  NOW(), 
  NOW()
);

INSERT INTO recipes (
  name, 
  description, 
  photo_url, 
  ingredients, 
  steps, 
  total_steps_time, 
  updated_at, 
  created_at
) VALUES (
  'Chicken Fajitas', 
  'Tex-Mex dish made with marinated chicken, bell peppers, onions, and spices.', 
  NULL, 
  '{chicken breasts, bell peppers, onions, chili powder, cumin}', 
  ARRAY[
    '{"title": "Marinate the chicken", "description": "In a large bowl, combine the chicken, chili powder, cumin, salt, and pepper. Toss well to coat the chicken with the spices. Cover and marinate in the refrigerator for at least 30 minutes or up to 2 hours."}'::jsonb, 
    '{"title": "Cook the chicken and vegetables", "description": "Heat a large skillet over medium-high heat. Add the chicken and cook until browned and cooked through. Remove the chicken from the skillet and set aside. Add the bell peppers and onions to the skillet and cook until softened."}'::jsonb, 
    '{"title": "Combine the chicken and vegetables", "description": "Return the chicken to the skillet and toss well to combine with the vegetables."}'::jsonb
  ], 
  '00:30:00', 
  NOW(), 
  NOW()
);