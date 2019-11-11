---
header-includes:
	\usepackage[czech]{babel}
	\usepackage{a4wide}
---
# Máš kliku

## Řešení

Proúčely řešení jsem si letiště a spoje představil jako neorientovaný ohodnocený graf, ohodnocení je cena cesty. Spoje 
mezi letišti v paktu jsem do něj nezahrnul, pouze jako 2 čísla: do jakého čísla jsou letiště v paktu a cenu za let. 
Následně jsem pustil hlednání nejkratších cest v ohodnoceném grafu z letiště Václava Havla (dijkstrův algoritmus) 
a zapsal si u každého vrcholu výslednou cenu. Letišť v paktu jsem z důvodu paměťové úspory při potkání 1 z nich prošel 
všechna i bez hran, abych je všechny nemusel ukládat.

### Složitost

Použil jsem dijkstrův algoritmus s haldou. Každý vrchol otevřu jednou, a v tomtu kroku projdeme všechny jeho následníky, 
úpravy haldy nám zaberou `log n`. Celková složitost bude tedy cca: `O(N*log N)` V prvním `N` je zahrnut počet vrcholů 
i počet jejich následníků.
