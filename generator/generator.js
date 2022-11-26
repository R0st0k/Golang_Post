import fs from 'fs'
import {v4 as uuidv4} from 'uuid';
import {default as _, isNull} from 'underscore'
import { ObjectID, UUID } from 'bson'

import Chance from 'chance'
let chance = new Chance();

const stageName = {
    REGISTERED: "Принято в отделении связи",
    LEFT_POST_OFFICE: "Покинуло место приема",
    ARRIVED_AT_POST_OFFICE: "Прибыло в место вручения",
    ARRIVED_AT_SORTING_OFFICE: "Прибыло в сортировочный центр",
    LEFT_SORTING_OFFICE: "Покинуло сортировочный центр",
    DELIVERED: "Вручено адресату",
}

const employeePosition = {
    POST_OFFICE_STAFF: "Сотрудник отделения связи",
    SORTING_OFFICE_STAFF: "Сотрудник сортировочного центра",
    DRIVER: "Водитель",
    POSTMAN: "Почтальон",
}

const officeType = {
    POST_OFFICE: "Отделение связи",
    SORTING_OFFICE: "Сортировочный центр",
}

const sendingsType = {
    LETTER: "Письмо",
    PARCEL: "Бандероль",
    PACKAGE: "Посылка"
}

const sendingStatus = {
    ON_THE_WAY: "В пути",
    LOST: "Потеряно",
    DELIVERED: "Доставлено"
}

function getOneOf(array) {
    return _.sample(array, 1)[0]
}

class PostGenerator {
    constructor() {
        this.employees = []
        this.post_offices = []
        this.postcodes = new Set()
        this.sendings = []
        this.counter = 1
    }

    #counter() {
        return this.counter++
    }

    #generatePostcodes(n) {
        let postcode_prefixes = [
            "450", "451", "452", "453",
            "420", "421", "422", "423",
            "190", "191", "192", "193", "194"
        ]

        for (; this.postcodes.size < n;) {
            let postcode_prefix = _.sample(postcode_prefixes, 1)[0]
            this.postcodes.add(postcode_prefix + chance.string({ length: 3, numeric: true }))
        }
    }

    #generateAddress(postcode = "", savePostcode = false) {
        function generateStreet() {
            return _.sample([
                "ул. ", "бул. ", "пр-кт ", "ал. "
            ], 1)[0] + _.sample([
                "Мира", "Кефира", "Эчпочмака",
                "Татаров", "Победы", "Силы",
                "Ленина", "ЛЭТИ", "Космонавтов",
                "Заки Валиди", "Габдуллы Тукая", "Мусы Джалиля",
                "Салавата Юлаева", "Мажита Гафури", "Губкина"
            ], 1)[0]
        }

        function generateBuilding() {
            return chance.integer({ min: 1, max: 100 }).toString()
        }

        let postcode_prefix = postcode.slice(0, 3)
        switch (postcode_prefix) {
            case "450":
            case "451":
            case "452":
            case "453":
                return {
                    postcode: (savePostcode) ? postcode : postcode_prefix + chance.string({ length: 3, numeric: true }),
                    region: "Республика Башкортостан",
                    district: _.sample([
                        "Стерлитамакский р-н",
                        "Мелеузосвский р-н",
                        "Уфимский р-н",
                        "Буздякский р-н",
                        "Туймазинский р-н",
                    ], 1)[0],
                    settlement: _.sample([
                        "г. Агидель",
                        "г. Кумертау",
                        "г. Межгорье",
                        "г. Нефтекамск",
                        "г. Октябрьский",
                        "г. Салават",
                        "г. Сибай",
                        "г. Стерлитамак",
                        "г. Уфа",
                        "п. Нугуш"
                    ], 1)[0],
                    street: generateStreet(),
                    building: generateBuilding()
                }
            case "420":
            case "421":
            case "422":
            case "423":
                return {
                    postcode: (savePostcode) ? postcode : postcode_prefix + chance.string({ length: 3, numeric: true }),
                    region: "Республика Татарстан",
                    district: _.sample([
                        "Азнакаевский р-н",
                        "Бугульминский р-н",
                        "Зеленодольский р-н",
                        "Менделеевский р-н",
                        "Нурлатский р-н",
                        "Тукаевский р-н",
                    ], 1)[0],
                    settlement: _.sample([
                        "г. Казань",
                        "г. Набережные Челны",
                        "г. Азнакаево",
                        "с. Мальбагуш",
                        "с. Сарлы",
                        "п. Тырыш",
                    ], 1)[0],
                    street: generateStreet(),
                    building: generateBuilding()
                }
            case "460":
            case "461":
            case "462":
                return {
                    postcode: (savePostcode) ? postcode : postcode_prefix + chance.string({ length: 3, numeric: true }),
                    region: "Оренбургаская область",
                    district: _.sample([
                        "Абдулинский р-н",
                        "Беляевский р-н",
                        "Оренбургский р-н",
                        "Грачевский р-н",
                        "Северный р-н",
                        "Тюльганский р-н",
                    ], 1)[0],
                    settlement: _.sample([
                        "г. Оренбург",
                        "г. Бузулук",
                        "г. Гай",
                        "г. Ясный",
                        "с. Булатовка",
                        "п. Искра",
                        "с. Лоховка",
                    ], 1)[0],
                    street: generateStreet(),
                    building: generateBuilding()
                }
            case "190":
            case "191":
            case "192":
            case "193":
            case "194":
                return {
                    postcode: (savePostcode) ? postcode : postcode_prefix + chance.string({ length: 3, numeric: true }),
                    region: "г. Санкт-Петербург",
                    settlement: _.sample([
                        "г. Санкт-Петербург",
                        "г. Зеленогорск",
                        "г. Петергоф",
                        "г. Пушкин",
                        "г. Ломоносов",
                        "г. Красное Село",
                        "г. Сестрорецк",
                    ], 1)[0],
                    street: generateStreet(),
                    building: generateBuilding()
                }
        }
    }

    #generateEmployees(postOfficeType) {
        let positions
        switch (postOfficeType) {
            case officeType.POST_OFFICE:
                positions = [
                    employeePosition.POST_OFFICE_STAFF,
                    employeePosition.POSTMAN,
                    employeePosition.DRIVER,
                ]
                break
            case officeType.SORTING_OFFICE:
                positions = [
                    employeePosition.SORTING_OFFICE_STAFF,
                    employeePosition.DRIVER,
                ]
                break
        }

        let employees_buffer = []
        for (let pos of positions) {
            let gender = _.sample(['male', 'female'], 1)[0]
            let haveMiddleName = chance.bool()
            let year = chance.year({ min: 1960, max: 2001 });
            let fullName = this.#generateRussianFullName(gender)
            let employee = {
                _id: new ObjectID(this.#counter()),
                surname: fullName.surname,
                name: fullName.name,
                gender: ((gender === 'male') ? 'М' : 'Ж'),
                birth_date: chance.birthday({ year: year }).toISOString(),
                position: pos,
                phone_number: "8" + chance.string({ length: 10, numeric: true }),
            }
            if (haveMiddleName) {
                employee.middle_name = fullName.middle_name
            }
            employees_buffer.push(employee)
        }

        return employees_buffer
    }

    #generatePostOffices(n) {
        if (n < 3) {
            throw "invalid number of post offices"
        }
        function generateOfficeType() {
            return _.sample([
                officeType.SORTING_OFFICE,
                officeType.POST_OFFICE
            ], 1)[0]
        }

        this.#generatePostcodes(n)
        let post_offices_buffer = []
        let employees_buffer = []
        for (let postcode of this.postcodes) {
            let officeType = generateOfficeType()
            let current_employees = this.#generateEmployees(officeType)
            employees_buffer.push(...current_employees)
            post_offices_buffer.push({
                _id: new ObjectID(this.#counter()),
                type: officeType,
                address: this.#generateAddress(postcode),
                employees: current_employees.map(x => x._id)
            })
        }

        // Нужно как минимум два отделения связи и один сортировочный центр
        if (
            post_offices_buffer.filter(x => x.type === officeType.POST_OFFICE).length > 1 &&
            post_offices_buffer.filter(x => x.type === officeType.SORTING_OFFICE).length > 0
        ) {
            this.post_offices.push(...post_offices_buffer)
            this.employees.push(...employees_buffer)
            return this.post_offices
        }

        throw "can't generate post offices, try again"
    }

    #generateSendings(n) {
        if (this.post_offices < 1) {
            throw "no post offices - no sendings"
        }
        function generateStagesAndStatus(sending, employees, post_offices) {
            function getEmployeeID(postcode, position) {
                let valid_office = _.sample(post_offices.filter(x => x.address.postcode === postcode), 1)[0]
                let employee = _.sample(employees.filter(x => {
                    return ((x.position === position) && (valid_office.employees.includes(x._id)));
                }), 1)[0]
                return employee._id
            }

            function addDays(iso_date, days) {
                let date = new Date(iso_date)
                date.setDate(date.getDate() + days)
                return date.toISOString()
            }

            let stages_buffer = [
                {
                    "name": stageName.REGISTERED,
                    "timestamp": sending.registration_date,
                    "postcode": sending.sender.address.postcode,
                    "employee_id": getEmployeeID(sending.sender.address.postcode, employeePosition.POST_OFFICE_STAFF),
                },
                {
                    "name": stageName.LEFT_POST_OFFICE,
                    "timestamp": addDays(sending.registration_date, 1),
                    "postcode": sending.sender.address.postcode,
                    "employee_id": getEmployeeID(sending.sender.address.postcode, employeePosition.POST_OFFICE_STAFF),
                }
            ]

            let sort_centers = post_offices.filter(x => x.type === officeType.SORTING_OFFICE)
            const max_intermediate_stages = 4
            for (let i = 0; i < max_intermediate_stages && chance.bool(); i++) {
                let last_stage = stages_buffer[stages_buffer.length - 1]
                let sort_center = _.sample(sort_centers, 1)[0]
                // Проверяем, что новый офис не встречался раньше
                let previous_postcodes = stages_buffer.map(x => x.postcode)
                if (previous_postcodes.includes(sort_center.address.postcode)) {
                    continue
                }
                stages_buffer.push({
                    "name": stageName.ARRIVED_AT_SORTING_OFFICE,
                    "timestamp": addDays(last_stage.timestamp, 1),
                    "postcode": sort_center.address.postcode,
                    "employee_id": getEmployeeID(sort_center.address.postcode, employeePosition.DRIVER),
                })
                stages_buffer.push({
                    "name": stageName.LEFT_SORTING_OFFICE,
                    "timestamp": addDays(last_stage.timestamp, 2),
                    "postcode": sort_center.address.postcode,
                    "employee_id": getEmployeeID(sort_center.address.postcode, employeePosition.SORTING_OFFICE_STAFF),
                })
            }

            // Если отправление было хотя бы раз в сортировочном центре, то пробуем его доставить или потерять
            if (stages_buffer.length > 2) {
                let isDelivered = chance.bool()
                if (isDelivered) {
                    let last_stage = stages_buffer[stages_buffer.length - 1]
                    stages_buffer.push({
                        "name": stageName.ARRIVED_AT_POST_OFFICE,
                        "timestamp": addDays(last_stage.timestamp, 1),
                        "postcode": sending.receiver.address.postcode,
                        "employee_id": getEmployeeID(sending.receiver.address.postcode, employeePosition.DRIVER),
                    })
                    stages_buffer.push({
                        "name": stageName.DELIVERED,
                        "timestamp": addDays(last_stage.timestamp, 2),
                        "postcode": sending.receiver.address.postcode,
                        "employee_id": getEmployeeID(sending.receiver.address.postcode, employeePosition.POSTMAN),
                    })
                    return {
                        stages: stages_buffer,
                        status: sendingStatus.DELIVERED
                    }
                }

                let isLost = chance.bool()
                if (isLost) {
                    return {
                        stages: stages_buffer,
                        status: sendingStatus.LOST
                    }
                }
            }

            return {
                stages: stages_buffer,
                status: sendingStatus.ON_THE_WAY
            }
        }

        let months = chance.month({ max: 10 });
        for (let i = 0; i < n; i++) {
            let senderFullName = this.#generateRussianFullName()
            let receiverFullName = this.#generateRussianFullName()
            let sending = {
                _id: new ObjectID(this.#counter()),
                order_id: new UUID(uuidv4()),
                registration_date: chance.date({ year: 2022, months: months }).toISOString(),
                sender: {
                    name: senderFullName.name,
                    surname: senderFullName.surname,
                    middle_name: senderFullName.middle_name,
                },
                receiver: {
                    name: receiverFullName.name,
                    surname: receiverFullName.surname,
                    middle_name: receiverFullName.middle_name,
                },
                type: _.sample([
                    sendingsType.LETTER,
                    sendingsType.PARCEL,
                    sendingsType.PACKAGE,
                ], 1)[0],
                size: {
                    length: chance.integer({ min: 10, max: 100}),
                    width: chance.integer({ min: 10, max: 100}),
                    height: chance.integer({ min: 10, max: 100}),
                },
                weight: chance.integer({ min: 10, max: 100}),
            }
            let valid_postcodes = this.post_offices.filter(x => x.type === officeType.POST_OFFICE).map(x => x.address.postcode)

            let senderPostcode = getOneOf(valid_postcodes)
            sending.sender.address = this.#generateAddress(senderPostcode, true)
            sending.sender.address.apartment = chance.integer({ min: 1, max: 100 }).toString()

            valid_postcodes = valid_postcodes.filter(postcode => postcode !== senderPostcode)
            sending.receiver.address = this.#generateAddress(getOneOf(valid_postcodes), true)
            sending.receiver.address.apartment = chance.integer({ min: 1, max: 100 }).toString()

            let stagesAndStatus = generateStagesAndStatus(sending, this.employees, this.post_offices)
            sending.stages = stagesAndStatus.stages
            sending.status = stagesAndStatus.status

            this.sendings.push(sending)
        }

        return this.sendings.map(x => x._id)
    }

    #generateRussianFullName(gender = '') {
        function generateMaleFullName() {
            const name = [
                "Руслан", "Ростислав", "Антон", "Даниил", "Данила", "Евгений",
                "Азат", "Александр", "Фёдор", "Артур", "Тимур", "Базослав",
                "Шайтан", "Рулон", "Венцеслав", "Аристарх", "Андрей", "Сергей", "Павел"
            ]
            const surname = [
                "Иванов", "Фамосов", "Бушанян", "Киранда", "Тапочек",
                "Обоев", "Гагарин", "Муссолини", "Сталин", "Бачини",
                "Кринжеборец", "Кеквейтов", "Пеперотов", "Трампов", "Бидонов"
            ]
            const middle_name = [
                "Русланович", "Ростиславович", "Антонович", "Базославович",
                "Рулонович", "Огрович", "Камжитосанович", "Бубльгумович",
                "Попович", "Карпович", "Мультифруктович", "Пиццович"
            ]
            return {
                name: getOneOf(name),
                surname: getOneOf(surname),
                middle_name: getOneOf(middle_name)
            }
        }

        function generateFemaleFullName() {
            const name = [
                "Камилла", "Гузель", "Элина", "Аделина", "Дарья", "Екатерина", "Ляйсан",
                "Фемочка", "Забава", "Илона", "Меган", "Бронислава", "Бажена", "Декабрина"
            ]
            const surname = [
                "Иванова", "Фамосова", "Бушанянова", "Кирандова", "Тапочек", "Бидонова",
                "Обоева", "Гагарина", "Кринжеборцева", "Кеквейтова", "Пеперотова", "Трампова"
            ]
            const middle_name = [
                "Руслановна", "Ростиславовна", "Антоновна", "Базославовна",
                "Рулоновна", "Огровна", "Камжитосановна", "Бубльгумовна",
                "Поповна", "Карповна", "Мультифруктовна", "Пиццовна"
            ]
            return {
                name: getOneOf(name),
                surname: getOneOf(surname),
                middle_name: getOneOf(middle_name)
            }
        }

        if (gender === '') {
            gender = 'male'
            if (chance.bool()) {
                gender = 'female'
            }
        }

        switch (gender) {
            case 'male':
                return generateMaleFullName()
            case 'female':
                return generateFemaleFullName()
            default:
                return "There is only two genders"
        }
    }

    gen(post_offices_cnt, sendings_cnt) {
        this.#generatePostOffices(post_offices_cnt)
        this.#generateSendings(sendings_cnt)
    }
}

const args = process.argv.slice(2)

let generator = new PostGenerator()
try {
    generator.gen(args[0], args[1])
} catch (e) {
    console.error(e)
    process.exit()
}

fs.writeFileSync('./samples/post_offices.json', JSON.stringify(generator.post_offices, null, 2), 'utf-8');
fs.writeFileSync('./samples/employees.json', JSON.stringify(generator.employees, null, 2), 'utf-8');
fs.writeFileSync('./samples/sendings.json', JSON.stringify(generator.sendings, null, 2), 'utf-8');