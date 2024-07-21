from .ChatAbstractServerMessage import ChatAbstractServerMessage
from .ChatServerMessage import ChatServerMessage
from .ExchangeTypesItemsExchangerDescriptionForUserMessage import ExchangeTypesItemsExchangerDescriptionForUserMessage
from .ExchangeTypesExchangerDescriptionForUserMessage import ExchangeTypesExchangerDescriptionForUserMessage
from .BasicPongMessage import BasicPongMessage

id_to_class = {
    1770: ChatAbstractServerMessage,
    6772: ChatServerMessage,
    2738: ExchangeTypesItemsExchangerDescriptionForUserMessage,
    6572: ExchangeTypesExchangerDescriptionForUserMessage,
    4877: BasicPongMessage
}