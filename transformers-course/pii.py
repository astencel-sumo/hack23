from transformers import pipeline

# pii = pipeline("ner",model="beki/en_spacy_pii_fast")
pii = pipeline(model="ArunaSaraswathy/bert-finetuned-ner-pii")
print(pii("I've been waiting for a HuggingFace course my whole life, and my name is Michael Bolton."))
print(pii("Weeks later, on December 7, Alareer was killed by a strike in Shajaiya, in northern Gaza, his friend and colleague, Jehad Abusalim, confirmed to CNN. He was staying with his brother, his sister, and her four children, who were also killed, according to Abusalim, a writer, 35, based in Washington, DC."))
